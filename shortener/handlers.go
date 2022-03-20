package shortener

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var shortener = Shortener{}

type App struct {
	Addr   string
	Router *mux.Router
	Server http.Server
}

func (app *App) Init() {
	app.Router = mux.NewRouter()
	app.Addr = ":" + os.Getenv("PORT")
	app.Server = http.Server{
		Addr:    app.Addr,
		Handler: app.Router,
	}
	shortener.Init()
	app.Router.HandleFunc("/", IndexHandler).Methods("GET")
	apiPathPrefix := app.Router.PathPrefix("/api").Subrouter()
	apiPathPrefix.HandleFunc("/shorten", ShortenHandler).Methods("POST")
	apiPathPrefix.HandleFunc("/getOrigin", GetOriginalHandler).Methods("POST")
	apiPathPrefix.HandleFunc("/shorten/custom", CustomShortenHandler).Methods("POST")
	apiPathPrefix.HandleFunc("/help", APIHelpPageHandler).Methods("GET")
	app.Router.HandleFunc("/s/{surl}", RedirectHandler).Methods("GET")
	app.Router.HandleFunc("/favicon.ico", faviconHandler)
	app.Router.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./static/js/"))))
	app.Router.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css/"))))
	app.Router.Use(RequestLoggerMiddleware(app.Router))
}

func (app *App) Run() {
	log.Println("App starting on ", app.Addr)
	app.Server.ListenAndServe()
	// log.Fatal(http.ListenAndServe(app.Addr, app.Router))
}

type HTMLPage struct {
	Title string
	Host  string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
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

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{
			"msg":   "Invalid request payload.",
			"error": err.Error(),
		})
		return
	}

	shortened, err := shortener.Shorten(body.URL)
	if err != nil {
		sendResponse(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		return
	}

	shorturl := r.Host + "/s/" + shortened.ShortUrl
	sendResponse(w, http.StatusOK, map[string]string{"URL": shorturl})
}

func GetOriginalHandler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload."})
		return
	}
	shortened, err := shortener.GetOriginalUrl(body.URL)
	if err != nil {
		sendResponse(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		return
	}

	originalurl := shortened.OriginUrl

	sendResponse(w, http.StatusOK, map[string]string{"URL": originalurl})
}

func CustomShortenHandler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload."})
		return
	}
	shortened, err := shortener.CustomShorten(body.URL, body.Custom)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	shorturl := r.Host + "/s/" + shortened.ShortUrl
	sendResponse(w, http.StatusOK, map[string]string{"URL": shorturl})
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	shorturl := mux.Vars(r)["surl"]

	var shortened Shortened
	shortener.DB.First(&shortened, "short_url = ?", shorturl)

	http.Redirect(w, r, shortened.OriginUrl, http.StatusSeeOther)

}

func APIHelpPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/apihelp.html")
	if err != nil {
		log.Println(err)
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	t.Execute(w, HTMLPage{Title: "GoShort / API", Host: r.Host})
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon/favicon.ico")
}
