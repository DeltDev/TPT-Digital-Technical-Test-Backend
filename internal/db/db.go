package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() *sqlx.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}