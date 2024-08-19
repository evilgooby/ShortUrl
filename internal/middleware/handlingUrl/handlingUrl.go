package handlingUrl

import (
	"ShortUrl/internal/middleware/generator"
	"ShortUrl/internal/repository/memory"
	"ShortUrl/internal/repository/postgres"
	"github.com/gin-gonic/gin"
)

func AddShortUrl(c *gin.Context, longUrl string, flagD bool) (string, error) {
	var shortUrl string
	shortUrl = VerifyShortUrl(longUrl, flagD)
	if shortUrl != "" {
		return shortUrl, nil
	}
	shortUrl, err := generator.GenerateShortLink(longUrl)
	if err != nil {
		return "", c.Error(err)
	}
	if !flagD {
		if err = postgres.AddInPostgres(shortUrl, longUrl); err != nil {
			return "", c.Error(err)
		}
	} else {
		memory.SaveUrl(shortUrl, longUrl)
	}
	return shortUrl, nil
}

func GetLongUrl(c *gin.Context, shortUrl string, flagD bool) (string, error) {
	var longUrl string
	if !flagD {
		url, err := postgres.GetLongUrlPostgres(shortUrl)
		if err != nil {
			return "", c.Error(err)
		}
		longUrl = url
	} else {
		url, err := memory.GetLongUrlMemory(shortUrl)
		if err != nil {
			return "", c.Error(err)
		}
		longUrl = url
	}
	return longUrl, nil
}

func VerifyShortUrl(longUrl string, flagD bool) string {
	if !flagD {
		str, _ := postgres.GetShortUrlPostgres(longUrl)
		if str != "" {
			return str
		}
	} else {
		str := memory.GetShortUrl(longUrl)
		if str != "" {
			return str
		}
	}
	return ""
}
