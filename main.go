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

    router.Run(":" + port)
}

func searchIncomeBySource(c *gin.Context) {
    source := c.Param("source")
    var matchingIncomes []Income
    for _, income := range incomeRecords {
        if income.Source == source {
            matchingIncomes = append(matchingIncomes, income)
        }
    }
    if len(matchingIncomes) > 0 {
        c.JSON(200, matchingIncomes)
    } else {
        c.JSON(404, gin.H{"message": "No incomes found for the provided source"})
    }
}

func searchExpenseByReason(c *gin.Context) {
    reason := c.Param("reason")
    var matchingExpenses []Expense
    for _, expense := range expenseRecords {
        if expense.Reason == reason {
            matchingExpenses = append(matchingExpenses, expense)
        }
    }
    if len(matchingExpenses) > 0 {
        c.JSON(200, matchingExpenses)
    } else {
        c.JSON(404, gin.H{"message": "No expenses found for the provided reason"})
    }
}

func generateComprehensiveReport(c *gin.Context) {
    var sourceList, reasonList []string
    sourceTotals := make(map[string]float64)
    reasonTotals := make(map[string]float64)

    for _, income := range incomeRecords {
        sourceTotals[income.Source] += income.Amount
        if !stringInSlice(sourceList, income.Source) {
            sourceList = append(sourceList, income.Source)
        }
    }

    for _, expense := range expenseRecords {
        reasonTotals[expense.Reason] += expense.Amount
        if !stringInSlice(reasonList, expense.Reason) {
            reasonList = append(reasonList, expense.Reason)
        }
    }

    sort.Strings(sourceList) // Alphabetize income sources
    sort.Strings(reasonList) // Alphabetize expense reasons

    c.JSON(200, gin.H{
        "total_income_sources":          len(sourceList),
        "income_source_totals":          sourceTotals,
        "total_expense_reasons":         len(reasonList),
        "expense_reason_totals":         reasonTotals,
        "alphabetized_income_sources":   sourceList,
        "alphabetized_expense_reasons":  reasonList,
    })
}

func stringInSlice(slice []string, str string) bool {
    for _, item := range slice {
        if item == str {
            return true
        }
    }
    return false
}