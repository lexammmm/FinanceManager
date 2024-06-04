package main

import (
    "fmt"
    "time"
)

type Income struct {
    Name   string
    Amount float64
    Date   time.Time
}

type Expense struct {
    Name   string
    Amount float64
    Date   time.Time
}

func logExpenses(expenses []float64, names []string, dates []time.Time) []Expense {
    var loggedExpenses []Expense

    for i, amount := range expenses {
        loggedExpenses = append(loggedExpenses, Expense{
            Amount: amount,
            Name:   names[i],
            Date:   dates[i],
        })
    }
    return loggedExpenses
}

func main() {
    var expenses = []float64{100.50, 200.75, 50.25}
    var names = []string{"Electricity", "Rent", "Internet"}

    // Optimizing the recurring computation of time.Now() by only calculating it once
    currentTime := time.Now()
    var dates = []time.Time{currentTime, currentTime.AddDate(0, 0, -30), currentTime.AddDate(0, 0, -60)}

    loggedExpenses := logExpenses(expenses, names, dates)

    for _, exp := range loggedExpenses {
        fmt.Printf("Logged Expense: %s - $%.2f - %s\n", exp.Name, exp.Amount, exp.Date.Format("2006-01-02"))
    }
}