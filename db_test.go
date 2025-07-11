package main

import (
	"database/sql"
	"finflow/config"
	"finflow/utils"
	"github.com/joho/godotenv"
	"testing"
)

func TestPostgresConnection(t *testing.T) {
	err := godotenv.Load()
	utils.PrintErr("Error loading .env file: ", err)

	connStr := config.GetDatabaseURL()
	db, err := sql.Open("postgres", connStr)
	utils.FatalErr("Could not connect to database", err)

	defer db.Close()

	utils.FatalErr("Error ping connection", db.Ping())

	t.Log("Successfully connected to the database")
}
