package main

import (
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