package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/batt0s/goshort/database"
	"github.com/batt0s/goshort/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type App struct {
	Addr     string
	Router   *chi.Mux
	Server   http.Server
	Database *database.Database

	ShortenerService services.ShortenerService
}

func (app *App) Init(appMode string) error {
	// Init Database
	database, err := database.New("sqlite", "dev.db", &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return err
	}
	app.Database = database

	// Init Services
	app.ShortenerService = services.NewShortenerService(app.Database.ShortenedRepo)

	// Init goth (package for Google auth)
	// InitGoth(appMode)

	// Init router
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST"},
	}))
	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))
	// Endpoints
	// Index Page
	r.Get("/", IndexHandler)
	// API Handlers
	r.Route("/api", func(api chi.Router) {
		// Shortener service API endpoints
		api.Route("/v3", func(sr chi.Router) {
			sr.Post("/shorten", app.ShortenHandler)
			// sr.Post("/customShorten", CustomShortenHandler)
			sr.Post("/getOrigin", app.GetOriginalHandler)
		})
	})
	// Short url redirect handler
	r.Get("/s/{shortUrl}", app.RedirectHandler)
	// other
	r.Get("/privacy", PrivacyHandler)
	// Static
	r.Handle("/static/js/*", http.StripPrefix("/static/js/", http.FileServer(http.Dir("./static/js/"))))
	r.Handle("/static/css/*", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css/"))))
	r.Handle("/static/img/*", http.StripPrefix("/static/img", http.FileServer(http.Dir("./static/img/"))))
	r.Get("/favicon.ico", faviconHandler)

	// Init HOST and PORT
	var host, port string
	host = os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
		log.Println("[warning] No HOST found, defaulting to 127.0.0.1")
	}
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("[warning] No PORT found, defaulting to 8080")
	}
	// Init app
	log.Println("App Mode : ", appMode)
	app.Router = r
	app.Addr = host + ":" + port
	app.Server = http.Server{
		Addr:    app.Addr,
		Handler: app.Router,
	}
	return nil
}

// Run App
func (app *App) Run() {
	log.Printf("[info] App starting on %s", app.Addr)
	app.Server.ListenAndServe()
	// log.Fatal(http.ListenAndServe(app.Addr, app.Router))
}
