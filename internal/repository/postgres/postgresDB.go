package postgres

import (
	"ShortUrl/config"
	"ShortUrl/internal/middleware/errorHandling"
	_ "github.com/lib/pq"
)

const (
	addDB       = `INSERT INTO shortened_urls (short_url, long_url)VALUES ($1, $2)`
	getLongUrl  = `SELECT long_url FROM shortened_urls WHERE short_url = $1`
	getShortUrl = `SELECT short_url FROM shortened_urls WHERE long_url = $1`
)

// Добавление данных в таблицу Postgres
func AddInPostgres(shortUrl string, longUrl string) error {
	_, err := config.StorePostgres.PostgresDB.Exec(addDB, shortUrl, longUrl)
	if err != nil {
		return errorHandling.ErrDB
	}
	return nil
}

// Получение длинной ссылки
func GetLongUrlPostgres(shortUrl string) (string, error) {
	var longUrl string
	if err := config.StorePostgres.PostgresDB.QueryRow(getLongUrl, shortUrl).Scan(&longUrl); err != nil {
		return "", errorHandling.ErrDB
	}
	return longUrl, nil
}

// Получение сокращенной ссылки
func GetShortUrlPostgres(longUrl string) (string, error) {
	var shortUrl string
	if err := config.StorePostgres.PostgresDB.QueryRow(getShortUrl, longUrl).Scan(&shortUrl); err != nil {
		return "", errorHandling.ErrDB
	}
	return shortUrl, nil
}
