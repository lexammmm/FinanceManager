package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	api := router.Group("/api")
	{
		income := api.Group("/incomes")
		{
			income.POST("/", createIncomeHandler)
			income.GET("/", getAllIncomesHandler)
			income.GET("/:id", getIncomeByIdHandler)
			income.PUT("/:id", updateIncomeByIdHandler)
			income.DELETE("/:id", deleteIncomeByIdHandler)
		}

		expenses := api.Group("/expenses")
		{
			expenses.POST("/", createExpenseHandler)
			expenses.GET("/", getAllExpensesHandler)
			expenses.GET("/:id", getExpenseByIdHandler)
			expenses.PUT("/:id", updateExpenseByIdHandler)
			expenses.DELETE("/:id", deleteExpenseByIdHandler)
		}
	}

	router.Run(":" + port)
}

func createIncomeHandler(c *gin.Context) {
	
}

func getAllIncomesHandler(c *gin.Context) {
	
}

func getIncomeByIdHandler(c *gin.Context) {
	
}

func updateIncomeByIdHandler(c *gin.Context) {
	
}

func deleteIncomeByIdHandler(c *gin.Context) {
	
}

func createExpenseHandler(c *gin.Context) {
	
}

func getAllExpensesHandler(c *gin.Context) {
	
}

func getExpenseByIdHandler(c *gin.Context) {
	
}

func updateExpenseByIdHandler(c *gin.Context) {
	
}

func deleteExpenseByIdHandler(c *gin.Context) {
	
}