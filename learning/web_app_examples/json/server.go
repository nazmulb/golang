package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nazmulb/golang/learning/web_app_examples/json/models"
	"gopkg.in/mgo.v2/bson"
)

var providers []models.Score

func getProviders(from string) ([]models.Score, error) {
	if from != "db" {
		file, _ := filepath.Abs("./providers.json")
		list, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(list, &providers)
		if err != nil {
			return nil, err
		}

	} else {
		session := models.GetDB()
		c := session.DB("myapp").C("scores")

		// Query All
		err := c.Find(bson.M{}).Sort("provider_user_id").All(&providers)
		if err != nil {
			return nil, err
		}
	}

	return providers, nil
}

// ProviderScores can be calculated
func ProviderScores(w http.ResponseWriter, req *http.Request) {
	qs := req.URL.Query()
	from := "db"
	if qs.Get("from") != "" {
		from = qs.Get("from")
	}

	providers, err := getProviders(from)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	fmt.Println(providers)

	var res []models.Score
	ids := []int{}
	m := make(map[int]bool) // hash to ensure uniq keys
	// Get all of the providers ids
	for _, id := range strings.Split(qs.Get("provider_user_ids"), ",") {
		n, err := strconv.Atoi(strings.TrimSpace(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		if _, ok := m[n]; !ok {
			m[n] = true
			ids = append(ids, n)
		}
	}

	for _, provider := range providers {
		for _, id := range ids {
			if id == provider.ProviderUserID {
				provider.CalculateScoreRank()
				provider.SetTypeOfWork(qs.Get("type_of_work_id"))
				res = append(res, provider)
			}
		}
	}

	js, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var data []models.Score
	err = json.Unmarshal([]byte(js), &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	tmpl := template.Must(template.ParseFiles("score.gtpl"))

	tmpl.Execute(w, data)
}

// Log server info
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	addr := os.Getenv("PROVIDER_MATCH_PORT")
	if addr == "" {
		addr = "9090"
	}

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", ProviderScores)

	log.Fatal(http.ListenAndServe(":"+addr, Log(http.DefaultServeMux)))
}
