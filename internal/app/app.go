package app

import (
	"ShortUrl/config"
	"ShortUrl/internal/controller"
	"flag"
	"log"
	"os"
)

func App() {
	config.ReadConfig()
	db := config.InitializeStorePostgres()
	defer db.PostgresDB.Close()
	config.NewLRUCache()
	dFlag := flag.Bool("d", false, "description of -d flag")
	flag.Parse()
	if *dFlag {
		err := os.Setenv("FLAG_D", "true")
		if err != nil {
			log.Fatalf("Failed to set flagD - Error: %v", err)
		}
	}
	controller.Registry()
}
