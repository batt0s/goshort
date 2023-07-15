package tests

import (
	"github.com/batt0s/goshort/database"
	"os"
	"testing"
)

var testCase = [2]string{"", "https://google.com"}
var testCaseCustom = [2]string{"battos", "https://batt0s.github.io"}

func TestMain(m *testing.M) {
	database.InitDB("test")
	exitVal := m.Run()
	os.Remove("test.db")
	os.Exit(exitVal)
}

func TestCreateShortened(t *testing.T) {
	shrtned, err := database.Shorten(testCase[1])
	if err != nil {
		t.Error(err)
	}
	if shrtned.OriginUrl != testCase[1] {
		t.Errorf("got %s, want %s", shrtned.OriginUrl, testCase[1])
	}
	if shrtned.ShortUrl == "" {
		t.Errorf("got short=''")
	}
	err = shrtned.Create()
	if err != nil {
		t.Error(err)
	}
	testCase[0] = shrtned.ShortUrl
}

func TestCreateShortenedCustom(t *testing.T) {
	shrtned, err := database.Shorten(testCaseCustom[1], testCaseCustom[0])
	if err != nil {
		t.Error(err)
	}
	if shrtned.OriginUrl != testCaseCustom[1] {
		t.Errorf("got %s, want %s", shrtned.OriginUrl, testCaseCustom[1])
	}
	if shrtned.ShortUrl != testCaseCustom[0] {
		t.Errorf("got %s, want %s", shrtned.ShortUrl, testCaseCustom[0])
	}
}

func TestGetOriginal(t *testing.T) {
	shrtned, err := database.GetOriginal(testCase[0])
	if err != nil {
		t.Error(err)
	}
	if shrtned.GetOrigin() != testCase[1] {
		t.Errorf("got %s, want %s", shrtned.GetOrigin(), testCase[1])
	}
}
