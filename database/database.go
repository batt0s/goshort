package database

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB

	ShortenedRepo ShortenedRepo
}

func New(driver string, source string, config *gorm.Config) (*Database, error) {
	if strings.TrimSpace(source) == "" {
		return nil, ErrorDatabaseSourceInvalid
	}
	db := &Database{}
	if err := db.connect(driver, source, config); err != nil {
		log.Println("Failed to connect database.")
		return nil, err
	}
	if err := db.db.AutoMigrate(&Shortened{}); err != nil {
		log.Println("Failed to migrate database.")
		return nil, err
	}
	db.ShortenedRepo = NewSqlShortenedRepo(db.db)
	return db, nil
}

func (db *Database) connect(driver string, source string, config *gorm.Config) error {
	if driver == "postgres" {
		return db._connect_postgres(source, config)
	} else if driver == "sqlite" {
		return db._connect_sqlite(source, config)
	}
	return ErrorDatabaseDriverInvalid
}

func (d *Database) _connect_sqlite(name string, config *gorm.Config) error {
	sqlDb := sqlite.Open(name)
	db, err := gorm.Open(sqlDb, config)
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *Database) _connect_postgres(dsn string, config *gorm.Config) error {
	sqlDb, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := sqlDb.PingContext(ctx); err != nil {
		return err
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDb,
	}), config)
	if err != nil {
		return err
	}
	d.db = db
	return nil
}
