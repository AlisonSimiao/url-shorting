package main

import (
	"log"
	API "url-shorting/api"
	db "url-shorting/database"
	routes "url-shorting/route"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	api := API.New()
	routes.Routes()

	db.Connect()

	api.Start()
}
