package main

import (
	"database/sql"
	"encoding/json"
	"finflow/models"
	"net/http"
)

func NewPaymentsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listPayments(w, db)
		case http.MethodPost:
			createPayment(w, r, db)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func listPayments(w http.ResponseWriter, db *sql.DB) {
	rows, err := db.Query("SELECT id, account_id, category_id, amount, direction, pay_date FROM payments ORDER BY id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var p models.Payment
		if err := rows.Scan(&p.ID, &p.AccountID, &p.CategoryID, &p.Amount, &p.Direction, &p.PayDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		payments = append(payments, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

func createPayment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var p models.Payment
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.QueryRow(
		"INSERT INTO payments(account_id, category_id, amount, direction, pay_date) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		p.AccountID, p.CategoryID, p.Amount, p.Direction, p.PayDate,
	).Scan(&p.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
