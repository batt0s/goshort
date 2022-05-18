package database

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type Shortened struct {
	gorm.Model
	OriginUrl string `json:"original_url"`
	ShortUrl  string `gorm:"unique" json:"short_url"`
	Author    string `json:"author"`
	Clicks    int    `json:"clicks"`
}

/**
* Shorten
* * Shorten the given URL
* @param myParam The parameter
* @param url The url to short
* @param author The one who shortened this url
* @param isCustom If user wants a custom short url
* @param custom The custom short url that user wants
 */
func Shorten(url string, author string, isCustom bool, custom string) (*Shortened, error) {
	// check if url is valid
	if !isValidURL(url) {
		return nil, fmt.Errorf("URL is not valid")
	}
	shortened := Shortened{}
	// check if there is a short url with the given url, if there is return
	if result := DB.Where("original_url = ?", url).First(&shortened); result.Error == nil {
		return &shortened, nil
	}
	var shortUrl string
	// check if short url is custom
	if isCustom {
		// check if the custom url already in database
		if _, err := GetOriginal(custom); err == nil {
			return nil, fmt.Errorf("custom url already in database")
		}
		shortUrl = custom
	} else {
		// generate a random short_url and check for if it is already in database
		for {
			shortUrl = generateRand(6)
			if _, err := GetOriginal(shortUrl); err != nil {
				break
			}
		}
	}
	shortened = Shortened{
		OriginUrl: url,
		ShortUrl:  shortUrl,
		Author:    author,
	}
	// insert to database
	if result := DB.Create(&shortened); result.Error != nil {
		return nil, result.Error
	}
	log.Println(shortened)
	return &shortened, nil
}

/**
* GetOriginal
* * Get original url of the given short url
* @param shortUrl The short url
 */
func GetOriginal(shortUrl string) (*Shortened, error) {
	shortened := &Shortened{}
	sp := strings.Split(shortUrl, "/")
	short := sp[len(sp)-1]
	if result := DB.Where("short_url = ?", short).First(shortened); result.Error != nil {
		return nil, fmt.Errorf("failed to get from database. error : %s", result.Error.Error())
	}
	return shortened, nil
}
