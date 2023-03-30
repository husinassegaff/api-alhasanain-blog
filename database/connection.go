package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	err := godotenv.Load("config/.env") // Load the .env file
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err.Error())
		os.Exit(1)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))

	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	DB = db
}
