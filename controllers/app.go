package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/batt0s/goshort/config"
	"github.com/batt0s/goshort/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type App struct {
	Addr   string
	Router *chi.Mux
	Server http.Server
}

func (app *App) Init(appMode string) {
	// Load config.json
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config files. Error : %s", err.Error())
	}
	// Init Database
	database.InitDB(appMode)
	// Init goth (package for google auth)
	InitGoth(appMode)
	// Init router
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST"},
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Endpoints
	r.Get("/", IndexHandler)
	// User auth service API endpoints
	r.Route("/auth", func(r chi.Router) {
		r.Get("/login", AuthHandler)
		r.Get("/callback", CallbackHandler)
		r.Get("/logout", HandleLogout)
	})
	r.Route("/api", func(api chi.Router) {
		// Shortener service API endpoints
		api.Route("/v2", func(sr chi.Router) {
			sr.Post("/shorten", ShortenHandler)
			// sr.Post("/customShorten", CustomShortenHandler)
			sr.Post("/getOrigin", GetOriginalHandler)
		})
	})
	// Short url redirect handler
	r.Get("/s/{shortUrl}", RedirectHandler)
	// other
	r.Get("/dashboard", DashboardHandler)
	r.Get("/privacy", PrivacyHandler)
	// Static
	r.Handle("/static/js/*", http.StripPrefix("/static/js/", http.FileServer(http.Dir("./static/js/"))))
	r.Handle("/static/css/*", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css/"))))
	r.Handle("/static/img/*", http.StripPrefix("/static/img", http.FileServer(http.Dir("./static/img/"))))
	r.Get("/favicon.ico", faviconHandler)

	// Init app
	log.Println("App Mode : ", appMode)
	var port string
	app.Router = r
	if appMode == "prod" {
		port = os.Getenv("PORT")
	} else {
		port = config.Conf.GetString(appMode + ".port")
	}
	if port == "" {
		port = "8080"
		log.Println("[warning] No port found")
	}
	host := config.Conf.GetString(appMode + ".host")
	app.Addr = host + ":" + port
	app.Server = http.Server{
		Addr:    app.Addr,
		Handler: app.Router,
	}

}

// Run App
func (app *App) Run() {
	log.Printf("[info] App starting on %s", app.Addr)
	app.Server.ListenAndServe()
	// log.Fatal(http.ListenAndServe(app.Addr, app.Router))
}
