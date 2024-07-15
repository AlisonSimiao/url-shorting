package main

import (
	"fmt"
	"log"
	db "vagas-api/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if db.Connect() != nil {
		fmt.Println("Conexão com o banco de dados realizada com sucesso!")
		return
	}

	//database := db.GetDatabase()

	fmt.Println("Init Migration")

	fmt.Println("Migration Success")
}
