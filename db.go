package main

import (
	"database/sql"
	"log"
	"os"
)

func InitDB() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "sfqwfw" {
		connStr := "postgres" // TODO some later
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