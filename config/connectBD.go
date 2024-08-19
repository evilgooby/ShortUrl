package config

import (
	"ShortUrl/internal/middleware/errorHandling"
	"database/sql"
	"github.com/spf13/viper"
	"log"
)

const (
	createTable = `CREATE TABLE shortened_urls (short_url TEXT PRIMARY KEY, long_url TEXT)`
	verifyTable = `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)`
)

type StoragePostgres struct {
	PostgresDB *sql.DB
}

var (
	StorePostgres = &StoragePostgres{}
)

// Подключение к базе данных
func InitializeStorePostgres() *StoragePostgres {
	connStr := viper.GetString("db.url")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	exists, err := TableExists(db, "shortened_urls")
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		_, err = db.Exec(createTable)
		if err != nil {
			log.Fatal(err)
		}
	}
	StorePostgres.PostgresDB = db
	return StorePostgres
}

// Проверка существует ли таблица
func TableExists(db *sql.DB, tableName string) (bool, error) {
	stmt, err := db.Prepare(verifyTable)
	if err != nil {
		return false, errorHandling.ErrDB
	}
	var exists bool
	if err = stmt.QueryRow(tableName).Scan(&exists); err != nil {
		return false, errorHandling.ErrDB
	}
	return exists, nil
}
