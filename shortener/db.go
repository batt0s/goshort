package shortener

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(mode string) *gorm.DB {
	var db *gorm.DB
	var err error
	if strings.ToLower(mode) != "prod" {
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
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
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
