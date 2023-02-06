package main


import (
	"fmt"
	"github.com/isabellabarcelos/url_shortener/handler"
	"net/http"
	"os"
	"errors"
	"encoding/json"
	"log"

)

func POST(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Create short URL\n")

	key:= r.URL.Query()

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["short_url"] = handler.CreateShortUrl(key.Get("long_url"))
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
	
}
func GET(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get URL\n")
	key := r.URL.Query()

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["long_url"] = handler.GetLink(key.Get("long_url"))
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return


}

func main() {
	http.HandleFunc("/", POST)
	http.HandleFunc("/get-url", GET)

	err := http.ListenAndServe(":8080", nil)

	
  if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
