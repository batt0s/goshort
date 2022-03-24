package shortener

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var shortener = Shortener{}

type App struct {
	Addr   string
	Router *chi.Mux
	Server http.Server
}

func (app *App) Init() {
	// Init custom logger
	CustomLogger.InitLogger()

	// Init shortener
	shortener.Init()

	// Init router

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Endpoints
	r.Get("/", IndexHandler)
	r.Route("/api/latest", func(sr chi.Router) {
		sr.Get("/docs", APIHelpPageHandler)
		sr.Post("/shorten", ShortenHandler)
		sr.Post("/customShorten", CustomShortenHandler)
		sr.Post("/getOrigin", GetOriginalHandler)
	})
	r.Get("/s/{shortUrl}", RedirectHandler)
	// Static
	r.Handle("/static/js/*", http.StripPrefix("/static/js/", http.FileServer(http.Dir("./static/js/"))))
	r.Handle("/static/css/*", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css/"))))
	r.Get("/favicon.ico", faviconHandler)

	// Init app
	app.Router = r
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		CustomLogger.Warn("Defaulting to PORT 8080")
	}
	app.Addr = ":" + port
	app.Server = http.Server{
		Addr:    app.Addr,
		Handler: app.Router,
	}

}

// Run App
func (app *App) Run() {
	CustomLogger.Info("App starting on ", app.Addr)
	app.Server.ListenAndServe()
	// log.Fatal(http.ListenAndServe(app.Addr, app.Router))
}

type HTMLPage struct {
	Title string
	Host  string
}

// Index Page Handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		CustomLogger.Error(err)
		//log.Println(err)
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	t.Execute(w, HTMLPage{Title: "GoShort", Host: r.Host})
}

// Request body
type RequestBody struct {
	URL    string `json:"url"`
	Custom string `json:"custom"`
}

// Function for Sending Response
func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// Shorten URL Handler
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{
			"msg":   "Invalid request payload.",
			"error": err.Error(),
		})
		CustomLogger.Error(err)
		return
	}

	shortened, err := shortener.Shorten(body.URL)
	if err != nil {
		sendResponse(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		CustomLogger.Error(err)
		return
	}

	shorturl := r.Host + "/s/" + shortened.ShortUrl
	sendResponse(w, http.StatusOK, map[string]string{"URL": shorturl})
}

// Handler : Getting Original URL from Short URL
func GetOriginalHandler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload."})
		CustomLogger.Error(err)
		return
	}
	shortened, err := shortener.GetOriginalUrl(body.URL)
	if err != nil {
		sendResponse(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		CustomLogger.Error(err)
		return
	}

	originalurl := shortened.OriginUrl

	sendResponse(w, http.StatusOK, map[string]string{"URL": originalurl})
}

// Handler : Shorten URL with a custom short url
func CustomShortenHandler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload."})
		CustomLogger.Error(err)
		return
	}
	shortened, err := shortener.CustomShorten(body.URL, body.Custom)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		CustomLogger.Error(err)
		return
	}
	shorturl := r.Host + "/s/" + shortened.ShortUrl
	sendResponse(w, http.StatusOK, map[string]string{"URL": shorturl})
}

// Redirect short url
func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	shorturl := chi.URLParam(r, "shortUrl")

	var shortened Shortened
	shortener.DB.First(&shortened, "short_url = ?", shorturl)

	http.Redirect(w, r, shortened.OriginUrl, http.StatusSeeOther)

}

// API Docs Page Handler
func APIHelpPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/apihelp.html")
	if err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		CustomLogger.Error(err)
		return
	}
	t.Execute(w, HTMLPage{Title: "GoShort / API", Host: r.Host})
}

// Handler : Serve favicon.ico
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon/favicon.ico")
}
