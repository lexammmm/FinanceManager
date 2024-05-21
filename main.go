package main

import (
	"github.com/gin-gonic/gin"
	"os"
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

	router.POST("/expenses", createExpense)
	router.GET("/expenses", getExpenses)
	router.GET("/expenses/:id", getExpense)
	router.PUT("/expenses/:id", updateExpense)
	router.DELETE("/expenses/:id", deleteExpense)

	router.Run(":" + port)
}

func createIncome(c *gin.Context) {
	var newIncome Income
	if err := c.BindJSON(&newIncome); err != nil {
		return
	}
	incomes = append(incomes, newIncome)
	c.JSON(201, newIncome)
}

func getIncomes(c *gin.Context) {
	c.JSON(200, incomes)
}

func getIncome(c *gin.Context) {
	id := c.Param("id")
	for _, a := range incomes {
		if a.ID == id {
			c.JSON(200, a)
			return
		}
	}
	c.JSON(404, gin.H{"message": "income not found"})
}

func updateIncome(c *gin.Context) {
	id := c.Param("id")
	var updatedIncome Income
	if err := c.BindJSON(&updatedIncome); err != nil {
		return
	}
	for i, a := range incomes {
		if a.ID == id {
			incomes[i] = updatedIncome
			c.JSON(200, updatedIncome)
			return
		}
	}
	c.JSON(404, gin.H{"message": "income not found"})
}

func deleteIncome(c *gin.Context) {
	id := c.Param("id")
	for i, a := range incomes {
		if a.ID == id {
			incomes = append(incomes[:i], incomes[i+1:]...)
			c.JSON(200, gin.H{"message": "income deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "income not found"})
}

func createExpense(c *gin.Context) {
	var newExpense Expense
	if err := c.BindJSON(&newExpense); err != nil {
		return
	}
	expenses = append(expenses, newExpense)
	c.JSON(201, newExpense)
}

func getExpenses(c *gin.Context) {
	c.JSON(200, expenses)
}

func getExpense(c *gin.Context) {
	id := c.Param("id")
	for _, a := range expenses {
		if a.ID == id {
			c.JSON(200, a)
			return
		}
	}
	c.JSON(404, gin.H{"message": "expense not found"})
}

func updateExpense(c *gin.Context) {
	id := c.Param("id")
	var updatedExpense Expense
	if err := c.BindJSON(&updatedExpense); err != nil {
		return
	}
	for i, a := range expenses {
		if a.ID == id {
			expenses[i] = updatedExpense
			c.JSON(200, updatedExpense)
			return
		}
	}
	c.JSON(404, gin.H{"message": "expense not found"})
}

func deleteExpense(c *gin.Context) {
	id := c.Param("id")
	for i, a := range expenses {
		if a.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			c.JSON(200, gin.H{"message": "expense deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "expense not found"})
}