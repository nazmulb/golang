package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Score of providers
type Score struct {
	ProviderUserID               int     `json:"provider_user_id"`
	TypeOfWorkID                 string  `json:"type_of_work_id"`
	ProviderScoreRank            float32 `json:"provider_score_rank"`
	ProviderQualityScoreRank     float32 `json:"provider_quality_score_rank"`
	ProviderPerformanceScoreRank float32 `json:"provider_performance_score_rank"`
}

func (s *Score) calculateScoreRank() {
	s.ProviderScoreRank = (s.ProviderQualityScoreRank + s.ProviderPerformanceScoreRank) / 2
}

func (s *Score) setTypeOfWork(tow string) {
	s.TypeOfWorkID = tow
}

var providers []Score

func getProviders() ([]Score, error) {
	if len(providers) == 0 {
		file, _ := filepath.Abs("./providers.json")
		list, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(list, &providers)
		if err != nil {
			return nil, err
		}
	}

	return providers, nil
}

// ProviderScores can be calculated
func ProviderScores(w http.ResponseWriter, req *http.Request) {
	qs := req.URL.Query()
	providers, err := getProviders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var res []Score
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
				provider.calculateScoreRank()
				provider.TypeOfWorkID = qs.Get("type_of_work_id")
				res = append(res, provider)
			}
		}
	}

	js, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var data []Score
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
