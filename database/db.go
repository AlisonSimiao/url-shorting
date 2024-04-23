package db

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

// get connection to database with env
func Connect() error {
	DB_PASS := os.Getenv("DB_PASS")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	db, err := gorm.Open(postgres.Open("user="+DB_USER+" password="+DB_PASS+" host="+DB_HOST+" port="+DB_PORT+" dbname="+DB_NAME), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatal("ERR> database connection failed %s", err)
		return err
	}

	database = db

	return nil
}

func ConnectCredential(user, password, host, port, dbname string) error {
	conn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", user, password, host, port, dbname)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fmt.Printf("ERR> database connection failed")
		return err
	}

	database = db

	return nil
}

func GetDatabase() *gorm.DB {
	return database
}
