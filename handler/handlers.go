package handler

import (
	"fmt"
	"github.com/isabellabarcelos/url_shortener/shortener"
	"net/http"
	"log"

)

var (
	LinkList = map[string]string{}
)



// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}



func CreateShortUrl(c string) string{
	_, ok := LinkList[c] 
	if ok {
		fmt.Println("Already have this link")
		return LinkList[c]
	} 
	LinkList[c] = shortener.GenerateShortLink(c)

	return LinkList[c]
}

// GetURL - Find link that matches the shortened link in the linkList
func GetLink(ShortURL string) string{
	var LongUrl string
	i:=0
	for key, value := range LinkList {
			if value == ShortURL{
				LongUrl = key
				Redirect(LongUrl)		
				return LongUrl 
		}
			i++
		}
	fmt.Println("The link doesn't exist.")
	
	return LongUrl
}

// Redirect
func Redirect(Link string) {
	http.Handle("/", http.RedirectHandler(Link, http.StatusMovedPermanently))
	log.Fatal(http.ListenAndServe(":8080", nil))
}