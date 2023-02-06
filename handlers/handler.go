package handler

import (
	"fmt"
	"github.com/isabellabarcelos/URL_SHORTENER/shortener"
	"net/http"
	"log"
)

var (
	LinkList map[string]string
)

// SaveUrlMapping - Add a link to the linkList and generate a shorter link
func CreateShortUrl(LongLink string, UserId string) {

	_, ok := LinkList[LongLink] 
	if ok {
		fmt.Println("Already have this link")
		return
	}
	
	LinkList[LongLink] = shortener.GenerateShortLink(LongLink, UserId)
	return
	
}

// GetURL - Find link that matches the shortened link in the linkList
func GetLink(ShortLink string) {
	i:=0
	for key, value := range LinkList {
			if value == ShortLink{
				Redirect(key)		
				return 
		}
			i++
		}
	fmt.Println("Already have this link")
	return 
}

// Redirect
func Redirect(Link string) {
	http.Handle("/", http.RedirectHandler(Link, http.StatusMovedPermanently))
	log.Fatal(http.ListenAndServe(":8080", nil))
}