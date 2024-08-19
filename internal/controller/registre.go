package controller

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

var FlagsD = &Flags{}

type Flags struct {
	D bool
}

func Registry() {
	r := gin.Default()
	flag.BoolVar(&FlagsD.D, "d", false, "Use memory storage")
	flag.Parse()
	r.POST("/CreateShortUrl", CreateShortUrl)
	r.GET("/GetLongUrl", HandleGetShortUrl)
	r.GET("/:RedirectShortUrl", HandleShortUrlRedirect)
	port := viper.GetString("port")
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Failed to start the web server - Error: %v", err)
	}
}
