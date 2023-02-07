package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/isabellabarcelos/url_shortener/repository"
)

// API rest - JSON
type ShortenedURL struct {
	OriginalURL  string `json:"original_url"`
	ShortenedURL string `json:"shortened_url"`
}

type APIError struct {
	Error string `json:"error"`
}

type Handler struct {
	repo repository.Repo
}

func main() {
	h := Handler{repo: repository.New()}

	http.HandleFunc("/", h.CreateUrl)
	http.HandleFunc("/short/", h.GetUrl)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (h Handler) CreateUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request")
	key, ok := r.URL.Query()["url"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		err := APIError{Error: fmt.Sprintf("Failed to add link %v", key)}
		fmt.Printf("error: %v\n", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	value, found := h.repo.Get(key[0])
	if found {
		w.WriteHeader(http.StatusOK)
		err := ShortenedURL{OriginalURL: key[0], ShortenedURL: value}
		json.NewEncoder(w).Encode(err)
		return
	}

	shortS, err := h.repo.Insert(key[0])
	if err != nil {
		return
	}

	short := ShortenedURL{
		OriginalURL:  key[0],
		ShortenedURL: fmt.Sprintf("http://localhost:8080/short/%s", shortS),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(short)
}

func (h Handler) GetUrl(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathArgs := strings.Split(path, "/")
	value, found := h.repo.Get(pathArgs[2])
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		err := APIError{Error: "Failed to redirect link"}
		fmt.Printf("error: %v\n", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	log.Printf("Redirected to: %s", value)
	http.Redirect(w, r, value, http.StatusPermanentRedirect)
}
