package database

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shortened struct {
	ID          string `gorm:"type:uuid;primary_key;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	OriginalUrl string         `gorm:"not null;size:128;;" json:"original_url"`
	ShortUrl    string         `gorm:"not null;size:128;unique;;" json:"short_url"`
}

type ShortenedRepo interface {
	Find(ctx context.Context, shortUrl string) (Shortened, error)
	FindByOriginal(ctx context.Context, url string) (Shortened, error)
	Add(ctx context.Context, shortened Shortened) error
	Delete(ctx context.Context, shortened Shortened) error
}

type SqlShortenedRepo struct {
	db *gorm.DB
}

func NewSqlShortenedRepo(db *gorm.DB) *SqlShortenedRepo {
	return &SqlShortenedRepo{
		db: db,
	}
}

func (repo SqlShortenedRepo) Find(ctx context.Context, shortUrl string) (Shortened, error) {
	select {
	case <-ctx.Done():
		return Shortened{}, ErrorOperationCanceled
	default:
		var shortened Shortened
		result := repo.db.Where("short_url = ?", shortUrl).First(&shortened)
		return shortened, result.Error
	}
}

func (repo SqlShortenedRepo) FindByOriginal(ctx context.Context, url string) (Shortened, error) {
	select {
	case <-ctx.Done():
		return Shortened{}, ErrorOperationCanceled
	default:
		var shortened Shortened
		result := repo.db.Where("original_url = ?", url).First(&shortened)
		return shortened, result.Error
	}
}

func (repo SqlShortenedRepo) Add(ctx context.Context, shortened Shortened) error {
	select {
	case <-ctx.Done():
		return ErrorOperationCanceled
	default:
		err := shortened.create(repo.db)
		return err
	}
}

func (repo SqlShortenedRepo) Delete(ctx context.Context, shortened Shortened) error {
	select {
	case <-ctx.Done():
		return ErrorOperationCanceled
	default:
		err := shortened.delete(repo.db)
		return err
	}
}

// String method, returns
// Shortened<ID:ShortUrl:OriginalUrl:CreatedAt>
func (s Shortened) String() string {
	return fmt.Sprintf("Shortened<%s:%s:%s:%v>", s.ID, s.ShortUrl, s.OriginalUrl, s.CreatedAt)
}

func isValidURL(testUrl string) bool {
	_, err := url.ParseRequestURI(testUrl)
	return err == nil
}

// IsValid Check if the model is valid
func (s Shortened) isValid() bool {
	if len(s.OriginalUrl) == 0 {
		return false
	}
	// Check if the given URL is valid
	return isValidURL(s.OriginalUrl)
}

// Create model in database
func (s *Shortened) create(db *gorm.DB) error {
	if s.ShortUrl == "" {
		return ErrorInvalidShortened
	}
	// Return error if model is not valid
	if !s.isValid() {
		return ErrorInvalidShortened
	}
	// Generate ID
	s.ID = uuid.New().String()
	// insert to database
	result := db.Create(&s)
	return result.Error
}

func (s *Shortened) delete(db *gorm.DB) error {
	result := db.Delete(s)
	return result.Error
}
