package services

import (
	"context"
	"log"
	"strings"

	"github.com/batt0s/goshort/database"
	"github.com/google/uuid"
)

type ShortenerService interface {
	Shorten(ctx context.Context, url string, custom string) (database.Shortened, error)
	GetOriginal(ctx context.Context, shortUrl string) (database.Shortened, error)
}

type shortener struct {
	repo database.ShortenedRepo
}

func NewShortenerService(repo database.ShortenedRepo) ShortenerService {
	return &shortener{
		repo: repo,
	}
}

func (s *shortener) Shorten(ctx context.Context, url string, custom string) (database.Shortened, error) {
	shortened := database.Shortened{}
	custom = strings.TrimSpace(custom)
	if custom != "" {
		if len(custom) < 4 || !onlyLetterOrDigit(custom) {
			log.Println(custom)
			return shortened, ErrorInvalidCustom
		}
		if obj, err := s.repo.Find(context.Background(), custom); err == nil {
			return obj, ErrorCustomAlreadyExists
		}
		shortened.ShortUrl = custom
	} else {
		if obj, err := s.repo.FindByOriginal(context.Background(), url); err == nil {
			return obj, nil
		}
		for {
			var shortUrl string = uuid.New().String()[:6]
			if _, err := s.repo.Find(context.Background(), shortUrl); err == nil {
				continue
			} else {
				shortened.ShortUrl = shortUrl
				break
			}
		}
	}
	shortened.OriginalUrl = url
	log.Println(shortened)
	err := s.repo.Add(ctx, shortened)
	return shortened, err
}

func (s *shortener) GetOriginal(ctx context.Context, shortUrl string) (database.Shortened, error) {
	return s.repo.Find(ctx, shortUrl)
}
