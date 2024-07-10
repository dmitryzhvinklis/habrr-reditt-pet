package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set in .env file")
	}

	DB, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
