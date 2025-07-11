package main

import (
	"finflow/utils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	utils.PrintErr("Error loading .env file: ", err)
	db := InitDB()
	defer db.Close()

	http.HandleFunc("/v1/cashflow", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetCashflowHandler(db, w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server started on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
