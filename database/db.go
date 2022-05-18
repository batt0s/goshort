package database

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/batt0s/goshort/config"
)

var DB *gorm.DB

func InitDB(mode string) error {
	var db *gorm.DB
	var err error
	mode = strings.ToLower(mode)
	switch mode {
	case "prod":
		dsn := os.Getenv("DATABASE_URL")
		sqlDb, err := sql.Open("postgres", dsn)
		if err != nil {
			return err
		}
		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDb,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err != nil {
			return err
		}
	case "dev":
		name := config.Conf.GetString("dev.db.name")
		db, err = gorm.Open(sqlite.Open(name), &gorm.Config{})
		if err != nil {
			return err
		}
	case "test":
		name := config.Conf.GetString("test.db.name")
		db, err = gorm.Open(sqlite.Open(name), &gorm.Config{})
		if err != nil {
			return err
		}
	default:
		db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
		if err != nil {
			return err
		}
	}
	db.AutoMigrate(&Shortened{})
	DB = db
	log.Println("[info] Connected to database.")
	return nil
}
