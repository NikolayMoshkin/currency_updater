package main

import (
	"currency_updater/internal/app"
	"currency_updater/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cnf, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	(app.NewApp(&cnf)).Run()
}
