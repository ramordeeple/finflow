package main

import "net/http"

func main() {
	db := InitDB()
	defer db.Close()

	http.HandleFunc("/v1/cashflow", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			
		}
	})
}