package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Transaction struct {
	ID     int     `json:"id"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

var transactions = make([]Transaction, 0)
var nextID = 1

func main() {
	http.HandleFunc("/transactions", transactionsHandler)
	http.HandleFunc("/transactions/", transactionHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(transactions)
	case http.MethodPost:
		var transaction Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		transaction.ID = nextID
		nextID++
		transactions = append(transactions, transaction)
		json.NewEncoder(w).Encode(transaction)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/transactions/"):])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, transaction := range transactions {
		if transaction.ID == id {
			switch r.Method {
			case http.MethodGet:
				json.NewEncoder(w).Encode(transaction)
				return
			case http.MethodPut:
				var updatedTransaction Transaction
				err := json.NewDecoder(r.Body).Decode(&updatedTransaction)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				updatedTransaction.ID = id
				transactions[i] = updatedTransaction
				json.NewEncoder(w).Encode(updatedTransaction)
				return
			case http.MethodDelete:
				transactions = append(transactions[:i], transactions[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
		}
	}

	http.Error(w, "Transaction not found", http.StatusNotFound)
}