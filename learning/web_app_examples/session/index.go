package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var appSession *sessions.Session

var authKey = []byte("something-very-secret") // Authorization Key for Secret

var store = sessions.NewCookieStore(authKey)

func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 1, // 1 hour
		HttpOnly: true,
	}
}

func getSession(r *http.Request) *sessions.Session {

	log.Println("session before get", appSession)

	if appSession != nil {
		return appSession
	}

	session, err := store.Get(r, "sess_app")
	appSession = session

	log.Println("session after get", session)
	if err != nil {
		panic(err)
	}

	return session
}

func destroySessionValue(w http.ResponseWriter, r *http.Request, key interface{}) {
	session := getSession(r)
	delete(session.Values, key)
	session.Options.MaxAge = -1
	session.Save(r, w)
}

func setSessionValue(w http.ResponseWriter, r *http.Request, key interface{}, value interface{}) {
	session := getSession(r)
	session.Values[key] = value
	fmt.Printf("set session with key %v and value %v\n", key, value)
	session.Save(r, w)
}

func getSessionValue(w http.ResponseWriter, r *http.Request, key interface{}) interface{} {
	session := getSession(r)
	valWithOutType := session.Values[key]
	log.Println("returned value: ", valWithOutType)

	return valWithOutType
}

func count(w http.ResponseWriter, r *http.Request) {
	//destroySessionValue(w, r, "countnum")

	ct := getSessionValue(w, r, "countnum")

	if ct == nil {
		setSessionValue(w, r, "countnum", 1)
	} else {
		setSessionValue(w, r, "countnum", (ct.(int) + 1))
	}

	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")

	t.Execute(w, getSessionValue(w, r, "countnum"))

}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", count).Methods("GET")
	http.Handle("/", rtr)
	log.Println("Listening...")
	err := http.ListenAndServe(":9090", http.DefaultServeMux)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
