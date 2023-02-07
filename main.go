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

type APIError struct {
	Error string `json:"error"`
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
	fmt.Println("request")
	key, ok := r.URL.Query()["url"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		err := APIError{Error: fmt.Sprintf("Failed to add link %v", key)}
		fmt.Printf("error: %v\n", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	value, found := linkList[key[0]]
	if found {
		w.WriteHeader(http.StatusOK)
		err := ShortenedURL{OriginalURL: key[0], ShortenedURL: value}
		json.NewEncoder(w).Encode(err)
		return
	}

	genString := fmt.Sprint(rand.Int63n(1000))
	linkList[genString] = key[0]

	short := ShortenedURL{
		OriginalURL:  key[0],
		ShortenedURL: fmt.Sprintf("http://localhost:8080/short/%s", genString),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(short)
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathArgs := strings.Split(path, "/")
	log.Printf("Redirected to: %s", linkList[pathArgs[2]])
	http.Redirect(w, r, linkList[pathArgs[2]], http.StatusPermanentRedirect)
}
