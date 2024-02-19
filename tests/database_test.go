package tests

import (
	"context"
	"os"
	"testing"

	"github.com/batt0s/goshort/database"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *database.Database

var testCase database.Shortened = database.Shortened{
	OriginalUrl: "https://google.com",
	ShortUrl:    "test11",
}

func TestMain(m *testing.M) {
	testdb, err := database.New("sqlite", "test.db", &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Failed to connect database.")
	}
	db = testdb
	exitVal := m.Run()
	os.Remove("test.db")
	os.Exit(exitVal)
}

func TestAddShortened(t *testing.T) {
	err := db.ShortenedRepo.Add(context.Background(), testCase)
	if err != nil {
		t.Error(err)
	}
}

func TestFindShortened(t *testing.T) {
	shortened, err := db.ShortenedRepo.Find(context.Background(), testCase.ShortUrl)
	if err != nil {
		t.Error(err)
	}
	if shortened.OriginalUrl != testCase.OriginalUrl {
		t.Errorf("got %s, want %s", shortened.OriginalUrl, testCase.OriginalUrl)
	}
	if shortened.ShortUrl != testCase.ShortUrl {
		t.Errorf("got %s, want %s", shortened.ShortUrl, testCase.ShortUrl)
	}
	testCase.ID = shortened.ID
}

func TestFindByOriginalShortened(t *testing.T) {
	shortened, err := db.ShortenedRepo.FindByOriginal(context.Background(), testCase.OriginalUrl)
	if err != nil {
		t.Error(err)
	}
	if shortened.OriginalUrl != testCase.OriginalUrl {
		t.Errorf("got %s, want %s", shortened.OriginalUrl, testCase.OriginalUrl)
	}
	if shortened.ShortUrl != testCase.ShortUrl {
		t.Errorf("got %s, want %s", shortened.ShortUrl, testCase.ShortUrl)
	}
}

func TestDeleteShortened(t *testing.T) {
	if err := db.ShortenedRepo.Delete(context.Background(), testCase); err != nil {
		t.Error(err)
	}
	if shouldNil, err := db.ShortenedRepo.Find(context.Background(), testCase.ShortUrl); err == nil {
		t.Errorf("got %v, want %v", shouldNil, nil)
	}
}
