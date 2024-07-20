package main

import (
	"log"
	API "vagas-api/api"
	db "vagas-api/database"
	routes "vagas-api/route"

	"github.com/joho/godotenv"
)
// @title Vagas API
// @version 0.1
// @description docs backend vagas SPA
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name authorization
// @license.name MIT License
// @license.url 
// @host localhost:8080
// @BasePath /
// @schemes http

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
