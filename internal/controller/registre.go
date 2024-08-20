package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

// Регистрация маршрутов
func Registry() {
	r := gin.Default()
	r.POST("/CreateShortUrl", CreateShortUrl)
	r.GET("/GetLongUrl", HandleGetShortUrl)
	r.GET("/:RedirectShortUrl", HandleShortUrlRedirect)
	port := viper.GetString("port")
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Failed to start the web server - Error: %v", err)
	}
}
