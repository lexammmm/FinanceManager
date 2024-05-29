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
    http.ListenAndServe(":"+serverPort, nil)
}

func handleTransactions(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(transactionRecords)
    case http.MethodPost:
        var newTransaction Transaction
        if err := json.NewDecoder(r.Body).Decode(&newTransaction); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        newTransaction.ID = nextTransactionID
        nextTransactionID++
        transactionRecords = append(transactionRecords, newTransaction)
        json.NewEncoder(w).Encode(newTransaction)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func handleTransactionByID(w http.ResponseWriter, r *http.Request) {
    transactionID, err := strconv.Atoi(r.URL.Path[len("/transactions/"):])
    if err != nil {
        http.Error(w, "Invalid Transaction ID", http.StatusBadRequest)
        return
    }

    for index, transaction := range transactionRecords {
        if transaction.ID == transactionID {
            switch r.Method {
            case http.MethodGet:
                json.NewEncoder(w).Encode(transaction)
                return
            case http.MethodPut:
                var updatedTransaction Transaction
                if err := json.NewDecoder(r.Body).Decode(&updatedTransaction); err != nil {
                    http.Error(w, err.Error(), http.StatusBadRequest)
                    return
                }
                updatedTransaction.ID = transactionID
                transactionRecords[index] = updatedTransaction
                json.NewEncoder(w).Encode(updatedTransaction)
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