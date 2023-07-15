package database

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type Shortened interface {
	String() string
	IsValid() bool
	Create() error
	GetOrigin() string
}

type shortened struct {
	gorm.Model
	OriginUrl string `json:"original_url"`
	ShortUrl  string `gorm:"unique" json:"short_url"`
}

func NewShortened() Shortened {
	return &shortened{}
}

// String method, returns
// Shortened<ID:ShortUrl:OriginalUrl:CreatedAt>
func (s *shortened) String() string {
	return fmt.Sprintf("Shortened<%d:%s:%s:%v>", s.ID, s.ShortUrl, s.OriginUrl, s.CreatedAt)
}

// IsValid Check if the model is valid
func (s *shortened) IsValid() bool {
	// Check if the given URL is valid
	if !isValidURL(s.OriginUrl) {
		return false
	}
	// Check if short url already in use
	// TODO Make another way than using GetOriginal function
	if _, err := GetOriginal(s.ShortUrl); err == nil {
		return false
	}
	return true
}

// Create model in database
func (s *shortened) Create() error {
	// Return error if model is not valid
	if !s.IsValid() {
		return errors.New("shortened not valid")
	}
	// insert to database
	// insert to database
	if result := DB.Create(&s); result.Error != nil {
		return result.Error
	}
	log.Println(s)
	return nil
}

// GetOrigin Get original url of shortened
func (s *shortened) GetOrigin() string {
	return s.OriginUrl
}

// Shorten Shorten the given url
func Shorten(url string, custom ...string) (*shortened, error) {
	// check if url is valid
	if !isValidURL(url) {
		return nil, fmt.Errorf("URL is not valid")
	}
	if len(custom) > 1 {
		return nil, errors.New("too many args")
	}
	shrt := new(shortened)
	// check if there is a short url with the given url, if there is return
	if result := DB.Where("original_url = ?", url).First(shrt); result.Error == nil {
		return shrt, nil
	}
	var customShort, shortUrl string
	// check if short url is custom
	if len(custom) == 1 {
		customShort = custom[0]
	}
	shrt.OriginUrl = url
	for {
		if strings.TrimSpace(customShort) != "" {
			shortUrl = customShort
		}
		if shortUrl == "" {
			shortUrl = generateRand(6)
		}
		shrt.ShortUrl = shortUrl
		if shrt.IsValid() {
			break
		}
	}
	return shrt, nil
}

// GetOriginal Get original url of the given short url
func GetOriginal(shortUrl string) (Shortened, error) {
	if shortUrl == "" {
		return nil, errors.New("given shortUrl is empty")
	}
	shrt := new(shortened)
	sp := strings.Split(shortUrl, "/")
	short := sp[len(sp)-1]
	if result := DB.Where("short_url = ?", short).First(shrt); result.Error != nil {
		return nil, fmt.Errorf("failed to get from database. error : %s", result.Error.Error())
	}
	return shrt, nil
}
