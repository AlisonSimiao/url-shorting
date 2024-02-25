package main

import (
	"log"
	API "time-wise/api"
	db "time-wise/database"
	routes "time-wise/route"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := API.New()
	routes.Routes()

	db.Connect()

	api.Start()
}
