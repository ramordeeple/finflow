package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func InitDB() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:123@localhost:5432/finflow?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database ping error: ", err)
	}

	return db
}
