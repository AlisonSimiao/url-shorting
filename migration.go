package main

import (
	"fmt"
	"log"
	db "url-shorting/database"
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

	database := db.GetDatabase()

	fmt.Println("Init Migration")

	// Alterar o tamanho da coluna password
	if err := database.Exec("ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(255)").Error; err != nil {
		fmt.Printf("Erro ao alterar o tamanho da coluna: %v", err)
	}

	fmt.Println("Migration Success")
}
