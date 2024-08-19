package controller

import (
	"ShortUrl/internal/middleware/errorHandling"
	"ShortUrl/internal/middleware/handlingUrl"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type UrlLongCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}
type UrlShortCreationRequest struct {
	ShortUrl string `json:"short_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlLongCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortUrl, err := handlingUrl.AddShortUrl(c, creationRequest.LongUrl, FlagsD.D)
	if err != nil {
		errorHandling.ErrorHandler(c)
		return
	}
	host := viper.GetString("url.lh")
	c.JSON(200, gin.H{
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("RedirectShortUrl")
	longUrl, err := handlingUrl.GetLongUrl(c, shortUrl, FlagsD.D)
	if err != nil {
		errorHandling.ErrorHandler(c)
		return
	}
	c.Redirect(302, longUrl)
}

func HandleGetShortUrl(c *gin.Context) {
	var creationRequest UrlShortCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	longUrl, err := handlingUrl.GetLongUrl(c, creationRequest.ShortUrl[22:], FlagsD.D)
	if err != nil {
		errorHandling.ErrorHandler(c)
		return
	}
	c.JSON(200, gin.H{
		"long_url": longUrl,
	})
}
