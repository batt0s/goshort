package shortener

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var shortener = Shortener{}

type App struct {
	Router *mux.Router
}

func (app *App) Init() {
	app.Router = mux.NewRouter()
	shortener.Init()
	app.Router.HandleFunc("/", app.Index).Methods("GET")
	app.Router.HandleFunc("/shorten", app.Shorten).Methods("POST")
	app.Router.HandleFunc("/{surl}", app.Redirect).Methods("GET")
	app.Router.HandleFunc("/getOrigin", app.GetOriginalUrl).Methods("POST")
	app.Router.HandleFunc("/shorten/custom", app.CustomShorten).Methods("POST")
	app.Router.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./static/js/"))))
	app.Router.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css/"))))
}

func (app *App) Run(port string) {
	log.Println("App starting on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}

type indexPage struct {
	Title string
}

func (app *App) Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	t.Execute(w, indexPage{Title: "GoShort"})
}

type Body struct {
	URL    string `json:"url"`
	Custom string `json:"custom"`
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func (app *App) Shorten(w http.ResponseWriter, r *http.Request) {
	var body Body

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload."})
		return
	}

	shortened, err := shortener.Shorten(body.URL)
	if err != nil {
		sendResponse(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		return
	}

	shorturl := r.Host + "/" + shortened.ShortUrl
	log.Println(shorturl)
	sendResponse(w, http.StatusOK, map[string]string{"URL": shorturl})
}

func (app *App) GetOriginalUrl(w http.ResponseWriter, r *http.Request) {
	var body Body

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

func (app *App) CustomShorten(w http.ResponseWriter, r *http.Request) {
	var body Body

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
	shorturl := r.Host + "/" + shortened.ShortUrl
	sendResponse(w, http.StatusOK, map[string]string{"URL": shorturl})
}

func (app *App) Redirect(w http.ResponseWriter, r *http.Request) {

	shorturl := mux.Vars(r)["surl"]

	var shortened Shortened
	shortener.DB.First(&shortened, "short_url = ?", shorturl)

	http.Redirect(w, r, shortened.OriginUrl, http.StatusSeeOther)

}
