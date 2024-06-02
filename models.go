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

func logExpenses(expense *Expense, expenses []float64, names []string, dates []time.Time) []Expense {
	var loggedExpenses []Expense
	for i, amount := range expenses {
		expense.Amount = amount
		expense.Name = names[i]
		expense.Date = dates[i]
		loggedExpenses = append(loggedExpenses, *expense)
	}
	return loggedExpenses
}

func main() {
	expense := &Expense{}
	expenses := []float64{100.50, 200.75, 50.25}
	names := []string{"Electricity", "Rent", "Internet"}
	dates := []time.Time{time.Now(), time.Now().AddDate(0, 0, -30), time.Now().AddDate(0, 0, -60)}

	loggedExpenses := logExpenses(expense, expenses, names, dates)

	for _, exp := range loggedExpenses {
		fmt.Printf("Logged Expense: %s - $%.2f - %s\n", exp.Name, exp.Amount, exp.Date.Format("2006-01-02"))
	}
}