package main

import (
	"database/sql"
	"encoding/json"
	"finflow/models"
	"net/http"
)

func GetCashflowHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	query := `
		select a.name, sum(case when e.is_debit then e.amount_minor else -e.amount_minor end) 
		/ 100.0 as amount, c.code as currency 
		from entries e
		join accounts a on e.account_id = a.id
		join currencies c on e.currency_id = a.id
		group by a.name, c.code
		order by a.name;`

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	var results []models.Cashflow
	for rows.Next() {
		var cf = models.Cashflow{}
		if err := rows.Scan(&cf.AccountName, &cf.Amount, &cf.Currency); err != nil {
			http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, cf)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
