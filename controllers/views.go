package controllers

import (
	"log"
	"net/http"
	"text/template"
)

// Index Page Handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "templates/index.html")
	t, err := template.ParseFiles("templates/index.html", "templates/_header.html", "templates/_footer.html")
	if err != nil {
		log.Println(err)
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	t.Execute(w, map[string]interface{}{"Title": "GoShort", "host": r.Host})
}

func PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/privacy.html", "templates/_header.html", "templates/_footer.html")
	if err != nil {
		log.Println(err)
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	t.Execute(w, nil)
}

// Handler : Serve favicon.ico
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon/favicon.ico")
}
