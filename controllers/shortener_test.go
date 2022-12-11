package controllers_test

import (
	"bytes"
	"encoding/json"
	"github.com/batt0s/goshort/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

var baseUrl = "/api/v2"

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

func TestShorten(t *testing.T) {

	jsonData := []byte(`{"url":"https://google.com"}`)
	req, _ := http.NewRequest("POST", baseUrl+"/shorten", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

}

func TestCustomShorten(t *testing.T) {

	jsonData := []byte(`{
		"url":"https://google.com",
		"is_custom": true,
		"custom":"test"
	}`)
	req, _ := http.NewRequest("POST", baseUrl+"/shorten", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	res := executeRequest(req)

	var m map[string]string
	json.Unmarshal(res.Body.Bytes(), &m)
	if res.Code == http.StatusOK {
		if val, ok := m["url"]; ok {
			if val != res.Result().Request.Host+"/s/test" {
				t.Errorf("Expected `url` key of the response to be set to `%s/test`, got %s", res.Result().Request.Host, val)
			}
		}
	} else if res.Code == http.StatusBadRequest {
		if val, ok := m["error"]; ok {
			if val != "test is already in database." {
				t.Errorf("Excepted `error` key of the response to be set to `test is already in database.`, got %s", m["error"])
			}
		}
	} else {
		t.Errorf("Expected response code %d or %d, got %d\n", http.StatusOK, http.StatusBadRequest, res.Code)
	}

}

func TestGetOrigin(t *testing.T) {

	jsonData := []byte(`{"url":"test"}`)
	req, _ := http.NewRequest("POST", baseUrl+"/getOrigin", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	if res.Code == http.StatusOK {
		var m map[string]string
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["URL"] != "https://google.com" {
			t.Errorf("Excepted `url` key of the response to be set to `https://google.com`, got %s", m["URL"])
		}
	}
}

func TestRedirect(t *testing.T) {
	req, _ := http.NewRequest("GET", "/s/test", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusSeeOther, res.Code)
}
