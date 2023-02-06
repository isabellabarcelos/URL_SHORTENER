package main


import (
	"fmt"
	"github.com/isabellabarcelos/url_shortener/handler"
	"net/http"
	"io"
	"os"
	"errors"

)

func POST(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Create short URL\n")

	key, _:= r.URL.Query()["long_url"]
	
	io.WriteString(w,handler.CreateShortUrl(key))

}
func GET(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get URL\n")
	key, _:= r.URL.Query()["long_url"]
	io.WriteString(w, handler.GetLink(key))
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
