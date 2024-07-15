package main

import (
	"log"
	API "vagas-api/api"
	db "vagas-api/database"
	routes "vagas-api/route"

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
