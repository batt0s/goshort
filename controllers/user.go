package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/batt0s/goshort/database"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var (
	secretKey = os.Getenv("SECRET_KEY")
	maxAge    = 86400 * 2
	store     = sessions.NewCookieStore([]byte(secretKey))
)

func InitGoth(appMode string) {
	var secure bool = false
	if appMode == "prod" {
		secure = true
	}
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = secure
	gothic.Store = store
	clientKey := os.Getenv("GOOGLE_ID")
	secret := os.Getenv("GOOGLE_SECRET")
	if !secure {
		log.Println("ClientKey : ", clientKey)
		log.Println("Secret : ", secret)
		log.Println("SecretKey : ", secretKey)
	}
	goth.UseProviders(
		google.New(clientKey, secret, "http://goshrt.herokuapp.com/auth/callback?provider=google", "profile"),
	)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error : %s", err.Error())
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	session, err := gothic.Store.New(r, "session")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error : %s", err.Error())
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	session.Values["userid"] = user.UserID
	session.Values["username"] = user.Name
	session.Save(r, w)
	w.Header().Set("Location", "/dashboard")
	w.WriteHeader(http.StatusSeeOther)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	err := gothic.Logout(w, r)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "Error")
	}
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, err := gothic.Store.Get(r, "session")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Please login first. \nError : %s", err.Error())
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	userid, ok := session.Values["userid"].(string)
	if !ok {
		log.Println("Failed to get userid")
		fmt.Fprint(w, "Please login")
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	username, ok := session.Values["username"].(string)
	if !ok {
		log.Println("Failed to get username")
		fmt.Fprint(w, "Please login")
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	shorteneds, err := database.GetShortenedsByAuthor(userid)
	if err != nil {
		shorteneds = []database.Shortened{}
	}
	t, err := template.ParseFiles("templates/dashboard.html", "templates/_header.html", "templates/_footer.html")
	if err != nil {
		sendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		log.Println(err)
		return
	}
	t.Execute(w, map[string]interface{}{"Username": username, "UserID": userid, "Shorteneds": shorteneds})
}
