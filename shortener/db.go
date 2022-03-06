package shortener

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dev bool) *gorm.DB {
	var db *gorm.DB
	var err error
	if dev {
		dsn := "host=localhost user=postgres password=password dbname=goshort port=5432 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	} else {
		dsn := os.Getenv("DATABASE_URL")
		sqlDB, err := sql.Open("postgres", dsn)
		if err != nil {
			panic(err)
		}
		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}
	return db
}

type Shortened struct {
	gorm.Model
	OriginUrl string
	ShortUrl  string `gorm:"unique"`
}
