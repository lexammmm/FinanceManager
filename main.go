package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
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

var incomeRecords []Income
var expenseRecords []Expense

func main() {
    router := gin.Default()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    router.POST("/incomes", addIncome)
    router.GET("/incomes", listIncomes)
    router.GET("/incomes/:id", findIncomeByID)
    router.PUT("/incomes/:id", modifyIncome)
    router.DELETE("/incomes/:id", removeIncome)
    router.GET("/incomes/search/:source", searchIncomeBySource)

    router.POST("/expenses", addExpense)
    router.GET("/expenses", listExpenses)
    router.GET("/expenses/:id", findExpenseByID)
    router.PUT("/expenses/:id", modifyExpense)
    router.DELETE("/expenses/:id", removeExpense)
    router.GET("/expenses/search/:reason", searchExpenseByReason)

    router.GET("/report", generateSummaryReport)
    router.GET("/detailed-report", generateComprehensiveReport)

    if err := router.Run(":" + port); err != nil {
        log.Fatalf("Failed to start the server: %v", err)
    }
}

func searchIncomeBySource(c *gin.Context) {
}

func searchExpenseByReason(c *gin.Context) {
}

func generateComprehensiveReport(c *gin.Context) {
}

func stringInSlice(slice []string, str string) bool {
}