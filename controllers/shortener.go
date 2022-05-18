package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/batt0s/goshort/database"
	"github.com/go-chi/chi/v5"
)

// GET /api/latest/shorten
// Shorten URL Handler
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	// get request body
	var body RequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{
			"msg":   "Invalid request payload.",
			"error": err.Error(),
		})
		log.Println(err)
		return
	}
	// shorten url
	shortened, err := database.Shorten(body.URL, body.Author, body.IsCustom, body.Custom)
	if err != nil {
		sendResponse(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		log.Println(err)
		return
	}
	// send response
	shorturl := r.Host + "/s/" + shortened.ShortUrl
	sendResponse(w, http.StatusOK, map[string]string{"URL": shorturl})
}

// Handler : Getting Original URL from Short URL
func GetOriginalHandler(w http.ResponseWriter, r *http.Request) {
	// get request body
	var body RequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload."})
		log.Println(err)
		return
	}
	// get original url
	shortened, err := database.GetOriginal(body.URL)
	if err != nil {
		sendResponse(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		log.Println(err)
		return
	}
	// send response
	originalurl := shortened.OriginUrl
	sendResponse(w, http.StatusOK, map[string]string{"URL": originalurl})
}

// Handler : Shorten URL with a custom short url
// func CustomShortenHandler(w http.ResponseWriter, r *http.Request) {
// 	var body RequestBody

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&body); err != nil {
// 		sendResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload."})
// 		log.Println(err)
// 		return
// 	}
// 	shortened, err := shortenerService.CustomShorten(body.URL, body.Custom)
// 	if err != nil {
// 		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
// 		log.Println(err)
// 		return
// 	}
// 	shorturl := r.Host + "/s/" + shortened.ShortUrl
// 	sendResponse(w, http.StatusOK, map[string]string{"URL": shorturl})
// }

// Redirect short url
func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	shorturl := chi.URLParam(r, "shortUrl")

	var shortened database.Shortened
	database.DB.First(&shortened, "short_url = ?", shorturl)
	shortened.Clicks++
	database.DB.Save(&shortened)

	http.Redirect(w, r, shortened.OriginUrl, http.StatusSeeOther)

}
