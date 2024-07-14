package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() error {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return fmt.Errorf("unable to parse DATABASE_URL: %w", err)
	}

	DB, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	return nil
}
