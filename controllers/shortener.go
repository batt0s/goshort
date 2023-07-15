package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/batt0s/goshort/database"
	"github.com/go-chi/chi/v5"
)

// ShortenHandler : Shorten URL
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	// get request body
	body, err := getRequestBody(w, r)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}
	// shorten url
	shortened, err := database.Shorten(body.URL, body.Custom)
	if err != nil {
		sendResponse(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		log.Println(err)
		return
	}
	// save to database
	err = shortened.Create()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"msg":   "Error while inserting to database.",
			"error": err.Error(),
		})
		log.Println(err)
		return
	}
	// send response
	shortUrl := r.Host + "/s/" + shortened.ShortUrl
	sendResponse(w, http.StatusOK, map[string]string{"URL": shortUrl})
}

// GetOriginalHandler : Getting Original URL from Short URL
func GetOriginalHandler(w http.ResponseWriter, r *http.Request) {
	// get request body
	body, err := getRequestBody(w, r)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
	originalUrl := shortened.GetOrigin()
	sendResponse(w, http.StatusOK, map[string]string{"URL": originalUrl})
}

// RedirectHandler Redirect short url
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortUrl")
	shortened, err := database.GetOriginal(shortUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, shortened.GetOrigin(), http.StatusSeeOther)
}
