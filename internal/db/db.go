package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() *sqlx.DB {
	dsn := "postgres://tpttechnicaltest:tpttechnicaltest@localhost:5432/tpttechnicaltest?sslmode=disable"

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}