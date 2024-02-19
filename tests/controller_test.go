package tests_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/batt0s/goshort/controllers"
)

var baseUrl = "/api/v3"

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app := controllers.App{}
	app.Init("test")
	app.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d\n", expected, actual)
	}
}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Remove("dev.db")
	os.Exit(exitVal)
}

func TestShorten(t *testing.T) {

	jsonData := []byte(`{
		"url":"https://google.com"
	}`)
	req, _ := http.NewRequest("POST", baseUrl+"/shorten", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

}

func TestCustomShorten(t *testing.T) {

	jsonData := []byte(`{
		"url":"https://google.com",
		"custom":"testtest"
	}`)
	req, _ := http.NewRequest("POST", baseUrl+"/shorten", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]string
	json.Unmarshal(res.Body.Bytes(), &m)
	if m["url"] != "/s/testtest" {
		t.Errorf("Expected `url` key of the response to be set to `/s/testtest`, got `%s`", m["url"])
	}

}

func TestGetOrigin(t *testing.T) {

	jsonData := []byte(`{
		"url":"testtest"
	}`)
	req, _ := http.NewRequest("POST", baseUrl+"/getOrigin", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	if res.Code == http.StatusOK {
		var m map[string]string
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["url"] != "https://google.com" {
			t.Errorf("Excepted `url` key of the response to be set to `https://google.com`, got `%s`", m["url"])
		}
	}
}

func TestRedirect(t *testing.T) {
	req, _ := http.NewRequest("GET", "/s/testtest", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusSeeOther, res.Code)
}
