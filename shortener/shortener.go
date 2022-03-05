package shortener

import (
	"fmt"
	"log"

	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Shortener struct {
	DB *gorm.DB
}

func (app *Shortener) Init() {
	app.DB = InitDB()
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
