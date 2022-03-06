package shortener

import (
	"fmt"
	"log"
	"strings"

	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Shortener struct {
	DB *gorm.DB
}

func (app *Shortener) Init() {
	app.DB = InitDB(false)
	app.DB.AutoMigrate(&Shortened{})
}

func (app *Shortener) Shorten(toShort string) (*Shortened, error) {

	if !isValidURL(toShort) {
		return &Shortened{}, fmt.Errorf("URL is not valid.")
	}

	shortenedTest := Shortened{}

	err := app.DB.Where("origin_url = ?", toShort).First(&shortenedTest)
	if err.Error == nil {
		return &shortenedTest, err.Error
	}

	shortUrl := generateRand(5)
	for {
		err = app.DB.Where("short_url = ?", shortUrl).First(&shortenedTest)
		log.Println(err.Error)
		if err.Error == nil {
			shortUrl = generateRand(5)
		} else {
			break
		}
	}

	shortened := Shortened{OriginUrl: toShort, ShortUrl: shortUrl}

	err = app.DB.Create(&shortened)
	log.Println(err)
	if err.Error != nil {
		return &shortened, err.Error
	}

	log.Println("shortener.Shorten : ", shortened.ShortUrl)

	return &shortened, nil

	// TODO: hata kontrolü eklenecek

}

func (app *Shortener) GetOriginalUrl(shortUrl string) (*Shortened, error) {

	if !isValidURL(shortUrl) {
		return &Shortened{}, fmt.Errorf("Invalid URL.")
	}

	shortened := &Shortened{}

	sp := strings.Split(shortUrl, "/")
	short_url := sp[len(sp)-1]

	err := app.DB.Where("short_url = ?", short_url).First(shortened)
	if err.Error != nil {
		return &Shortened{}, fmt.Errorf("Couldn't found url in database. %v", err.Error)
	}

	return shortened, nil

}

func (app *Shortener) CustomShorten(toShort string, customUrl string) (*Shortened, error) {

	if !isValidURL(toShort) {
		return &Shortened{}, fmt.Errorf("Invalid URL.")
	}

	shortened := &Shortened{}

	err := app.DB.Where("short_url = ?", customUrl).First(shortened)
	if err.Error == nil {
		return &Shortened{}, fmt.Errorf("%v is already in database.", customUrl)
	} else {
		shortened = &Shortened{OriginUrl: toShort, ShortUrl: customUrl}
		err = app.DB.Create(shortened)
		if err != nil {
			return shortened, err.Error
		}
	}

	return shortened, nil

}
