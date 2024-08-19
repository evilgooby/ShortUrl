package memory

import (
	"ShortUrl/internal/middleware/errorHandling"
)

var (
	storeMemory = &StorageMemory{
		Data: make(map[string]string),
	}
)

type StorageMemory struct {
	Data map[string]string
}

func SaveUrl(shortUrl string, originalUrl string) {
	storeMemory.Data[shortUrl] = originalUrl
}

func GetLongUrlMemory(shortUrl string) (string, error) {
	val, ok := storeMemory.Data[shortUrl]
	if !ok {
		return "", errorHandling.ErrNotFoundUrl
	}
	return val, nil
}

func GetShortUrl(longUrl string) string {
	for k, v := range storeMemory.Data {
		if v == longUrl {
			return k
		}
	}
	return ""
}
