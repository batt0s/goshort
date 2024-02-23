package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// ShortenHandler : Shorten URL
func (app *App) ShortenHandler(w http.ResponseWriter, r *http.Request) {
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
	shortened, err := app.ShortenerService.Shorten(context.Background(), body.URL, body.Custom)
	if err != nil {
		response := map[string]string{"error": err.Error()}
		sendResponse(w, http.StatusInternalServerError, response)
		log.Println(err)
		return
	}
	// send response
	shortUrl := "/s/" + shortened.ShortUrl
	response := map[string]string{"url": shortUrl}
	sendResponse(w, http.StatusOK, response)
}

// GetOriginalHandler : Getting Original URL from Short URL
func (app *App) GetOriginalHandler(w http.ResponseWriter, r *http.Request) {
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
	shortened, err := app.ShortenerService.GetOriginal(context.Background(), body.URL)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendResponse(w, http.StatusNotFound, nil)
		} else {
			sendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		log.Println(err)
		return
	}
	// send response
	originalUrl := shortened.OriginalUrl
	sendResponse(w, http.StatusOK, map[string]string{"url": originalUrl})
}

// RedirectHandler Redirect short url
func (app *App) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortUrl")
	shortened, err := app.ShortenerService.GetOriginal(context.Background(), shortUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, shortened.OriginalUrl, http.StatusSeeOther)
}
