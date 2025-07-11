package main

import (
	"database/sql"
	"finflow/config"
	"finflow/utils"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() *sql.DB {
	connStr := config.GetDatabaseURL()
	db, err := sql.Open("postgres", connStr)

	utils.FatalErr("Database connection error:", err)

	utils.FatalErr("Database ping error:", db.Ping())

	log.Println("Successfully connected to the database")
	return db
}
