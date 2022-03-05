package shortener

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var db *gorm.DB
	var err error
	dsn := "host=localhost user=postgres password=password dbname=goshort port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

type Shortened struct {
	gorm.Model
	OriginUrl string
	ShortUrl  string `gorm:"unique"`
}
