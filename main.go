package main


import (
	"fmt"
	"github.com/isabellabarcelos/url_shortener/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/shortUrl", func(c *gin.Context) {
		handler.GetLink(c)
	})


	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
