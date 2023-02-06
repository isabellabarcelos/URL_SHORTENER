package main

import (
	"fmt"
	"strings"
	"github.com/isabellabarcelos/url-shortener-2/handler"
	"github.com/isabellabarcelos/url-shortener-2/store"
)

var (
	LinkList map[string]string
)

// SaveUrlMapping - Add a link to the linkList and generate a shorter link
func SaveUrlMapping(LongLink string) {

	_, ok := LinkList[LongLink] 
	if ok {
		fmt.Println("Already have this link")
		return
	}
	
	LinkList[LongLink] = handler.GenerateUrl(LongLink)	
	return
	
}

// GetURL - Find link that matches the shortened link in the linkList
func getLink(ShortLink string) {
	i:=0
	for key, value := range LinkList {
			if value == ShortLink{
			    handler.redirect(key)		
				return 
		}
			i++
		}
	fmt.Println("Already have this link")
	return 
}
	

													

	
