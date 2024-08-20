package memory

import (
	"ShortUrl/config"
	"ShortUrl/internal/middleware/errorHandling"
)

// Сохранение данных в кэш
func SaveUrl(shortUrl string, longUrl string) {
	config.StoreMemory.Cache.Add(shortUrl, longUrl)
}

// Получение длинной ссылки из кэша
func GetLongUrlMemory(shortUrl string) (string, error) {
	val, ok := config.StoreMemory.Cache.Get(shortUrl)
	if !ok {
		return "", errorHandling.ErrNotFoundUrl
	}
	return val.(string), nil
}

// Получение короткой ссылки из кэша
func GetShortUrl(longUrl string) (string, error) {
	keys := config.StoreMemory.Cache.Keys()
	for _, key := range keys {
		val, ok := config.StoreMemory.Cache.Get(key)
		if ok {
			if val.(string) == longUrl {
				return val.(string), nil
			}
		}
	}
	return "", errorHandling.ErrNotFoundUrl
}
