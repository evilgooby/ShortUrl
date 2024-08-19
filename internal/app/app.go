package app

import (
	"ShortUrl/config"
	"ShortUrl/internal/controller"
)

func App() {
	config.ReadConfig()
	db := config.InitializeStorePostgres()
	defer db.PostgresDB.Close()
	controller.Registry()
}
