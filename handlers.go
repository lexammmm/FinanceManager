package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "os"
    "strconv"
    "strings"
)

type Transaction struct {
    ID     int     `json:"id"`
    Type   string  `json:"type"`
    Amount float64 `json:"amount"`
}

var transactionRecords = make([]Transaction, 0)
var nextTransactionID = 1

func main() {
    http.HandleFunc("/transactions", handleTransactions)
    http.HandleFunc("/transactions/", handleTransactionByID)

    serverPort := os.Getenv("PORT")
    if serverPort == "" {
        serverPort = "8080"
    }

    fmt.Printf("Server running on port %s\n", serverPort)
    err := http.ListenAndServe(":"+serverPort, nil)
    if err != nil {
        fmt.Println("Failed to start server:", err)
        os.Exit(1)
    }
}

func handleTransactions(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(transactionRecords); err != nil {
            http.Error(w, "Failed to encode transactions", http.StatusInternalServerError)
        }
    case http.MethodPost:
        var newTransaction Transaction
        if err := json.NewDecoder(r.Body).Decode(&newTransaction); err != nil {
            if errors.Is(err, io.EOF) {
                http.Error(w, "Empty request body", http.StatusBadRequest)
                return
            }
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }
        newTransaction.ID = nextTransactionID
        nextTransactionID++
        transactionRecords = append(transactionRecords, newTransaction)

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        if err := json.NewEncoder(w).Encode(newTransaction); err != nil {
            http.Error(w, "Failed to encode new transaction", http.StatusInternalServerError)
        }
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func handleTransactionByID(w http.ResponseWriter, r *http.Request) {
    transactionID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/transactions/"))
    if err != nil {
        http.Error(w, "Invalid Transaction ID", http.StatusBadRequest)
        return
    }

    for index, transaction := range transactionRecords {
        if transaction.ID == transactionID {
            switch r.Method {
            case http.MethodGet:
                w.Header().Set("Content-Type", "application/json")
                if err := json.NewEncoder(w).Encode(transaction); err != nil {
                    http.Error(w, "Failed to encode transaction", http.StatusInternalServerError)
                }
                return
            case http.MethodPut:
                var updatedTransaction Transaction
                if err := json.NewDecoder(r.Body).Decode(&updatedTransaction); err != nil {
                    http.Error(w, "Invalid request body", http.StatusBadRequest)
                    return
                }
                updatedTransaction.ID = transactionID
                transactionRecords[index] = updatedTransaction
                w.Header().Set("Content-Type", "application/json")
                if err := json.NewEncoder(w).Encode(updatedTransaction); err != nil {
                    http.Error(w, "Failed to encode updated transaction", http.StatusInternalServerError)
                }
                return
            case http.MethodDelete:
                transactionRecords = append(transactionRecords[:index], transactionRecords[index+1:]...)
                w.WriteHeader(http.StatusNoContent)
                return
            default:
                w.WriteHeader(http.StatusMethodNotAllowed)
                return
            }
        }
    }

    http.Error(w, "Transaction Not Found", http.StatusNotFound)
}