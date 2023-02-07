package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// API rest - JSON
type ShortenedURL struct {
	OriginalURL  string `json:"original_url"`
	ShortenedURL string `json:"shortened_url"`
}

var (
	linkList map[string]string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	linkList = map[string]string{}

	http.HandleFunc("/", CreateUrl)
	http.HandleFunc("/short/", GetUrl)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()["url"]
	if ok {
		if _, ok := linkList[key[0]]; !ok {
			genString := fmt.Sprint(rand.Int63n(1000))
			linkList[genString] = key[0]
			short := ShortenedURL{
				OriginalURL:  key[0],
				ShortenedURL: fmt.Sprintf("http://localhost:8080/short/%s", genString),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(short); err != nil {
				fmt.Printf("error: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Already have this link")
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Failed to add link")
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathArgs := strings.Split(path, "/")
	log.Printf("Redirected to: %s", linkList[pathArgs[2]])
	http.Redirect(w, r, linkList[pathArgs[2]], http.StatusPermanentRedirect)
}
