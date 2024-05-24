package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"sort"
)

type Income struct {
	ID     string  `json:"id"`
	Source string  `json:"source"`
	Amount float64 `json:"amount"`
}

type Expense struct {
	ID     string  `json:"id"`
	Reason string  `json:"reason"`
	Amount float64 `json:"amount"`
}

var incomes []Income
var expenses []Expense

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.POST("/incomes", createIncome)
	router.GET("/incomes", getIncomes)
	router.GET("/incomes/:id", getIncome)
	router.PUT("/incomes/:id", updateIncome)
	router.DELETE("/incomes/:id", deleteIncome)
	router.GET("/incomes/search/:source", searchIncomesBySource) // New route for searching incomes by source

	router.POST("/expenses", createExpense)
	router.GET("/expenses", getExpenses)
	router.GET("/expenses/:id", getExpense)
	router.PUT("/expenses/:id", updateExpense)
	router.DELETE("/expenses/:id", deleteExpense)
	router.GET("/expenses/search/:reason", searchExpensesByReason) // New route for searching expenses by reason

	router.GET("/report", generateReport)
	router.GET("/detailed-report", generateDetailedReport) // New route for detailed report generation

	router.Run(":" + port)
}

// Existing functions remain unchanged...

// New function to search incomes by source
func searchIncomesBySource(c *gin.Context) {
	source := c.Param("source")
	var filteredIncomes []Income
	for _, income := range incomes {
		if income.Source == source {
			filteredIncomes = append(filteredIncomes, income)
		}
	}
	if len(filteredIncomes) > 0 {
		c.JSON(200, filteredIncomes)
	} else {
		c.JSON(404, gin.H{"message": "No incomes found for the given source"})
	}
}

// New function to search expenses by reason
func searchExpensesByReason(c *gin.Context) {
	reason := c.Param("reason")
	var filteredExpenses []Expense
	for _, expense := range expenses {
		if expense.Reason == reason {
			filteredExpenses = append(filteredExpenses, expense)
		}
	}
	if len(filteredExpenses) > 0 {
		c.JSON(200, filteredExpenses)
	} else {
		c.JSON(404, gin.H{"message": "No expenses found for the given reason"})
	}
}

// New function to generate a detailed report
func generateDetailedReport(c *gin.Context) {
	var incomeSources, expenseReasons []string
	incomeMap := make(map[string]float64)
	expenseMap := make(map[string]float64)

	for _, income := range incomes {
		incomeMap[income.Source] += income.Amount
		if !contains(incomeSources, income.Source) {
			incomeSources = append(incomeSources, income.Source)
		}
	}

	for _, expense := range expenses {
		expenseMap[expense.Reason] += expense.Amount
		if !contains(expenseReasons, expense.Reason) {
			expenseReasons = append(expenseReasons, expense.Reason)
		}
	}

	sort.Strings(incomeSources) // Sort sources alphabetically
	sort.Strings(expenseReasons) // Sort reasons alphabetically

	c.JSON(200, gin.H{
		"total_income_sources":           len(incomeSources),
		"income_sources_with_amounts":    incomeMap,
		"total_expense_reasons":          len(expenseReasons),
		"expense_reasons_with_amounts":   expenseMap,
		"sorted_income_sources":          incomeSources,
		"sorted_expense_reasons":         expenseReasons,
	})
}

// Helper function to check if slice contains a string
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}