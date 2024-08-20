package handlingUrl

import (
	"ShortUrl/internal/middleware/generator"
	"ShortUrl/internal/repository/memory"
	"ShortUrl/internal/repository/postgres"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

// Генерируем и обрабатываем сокращенную ссылку
func AddShortUrl(c *gin.Context, longUrl string) (string, error) {
	var shortUrl string
	host := viper.GetString("url.lh")
	flagD := os.Getenv("FLAG_D")
	shortUrl, err := VerifyShortUrl(longUrl)
	if err != nil {
		log.Println("URL not found in database")
	}
	if shortUrl != "" {
		return host + shortUrl, nil
	}
	shortUrl, err = generator.GenerateShortLink(longUrl)
	if err != nil {
		return "", c.Error(err)
	}
	if flagD != "true" {
		if err = postgres.AddInPostgres(shortUrl, longUrl); err != nil {
			return "", c.Error(err)
		}
	} else {
		memory.SaveUrl(shortUrl, longUrl)
	}

	return host + shortUrl, nil
}

// Возвращаем длинную ссылку из базы данных или памяти
func GetLongUrl(c *gin.Context, shortUrl string) (string, error) {
	flagD := os.Getenv("FLAG_D")
	if flagD != "true" {
		url, err := postgres.GetLongUrlPostgres(shortUrl)
		if err != nil {
			return "", c.Error(err)
		}
		return url, nil
	} else {
		url, err := memory.GetLongUrlMemory(shortUrl)
		if err != nil {
			return "", c.Error(err)
		}
		return url, nil
	}
}

// Проверка есть ли в базе данных ссылка. Если да возвращаем её
func VerifyShortUrl(longUrl string) (string, error) {
	flagD := os.Getenv("FLAG_D")
	if flagD != "true" {
		str, err := postgres.GetShortUrlPostgres(longUrl)
		if str != "" {
			return str, nil
		}
		if err != nil {
			return "", err
		}
	} else {
		str, err := memory.GetShortUrl(longUrl)
		if str != "" {
			return str, nil
		}
		if err != nil {
			return "", err
		}
	}
	return "", nil
}
