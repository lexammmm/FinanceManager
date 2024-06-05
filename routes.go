package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()

	port := getServerPort()

	api := router.Group("/api")
	{
		setUpIncomeRoutes(api.Group("/incomes"))
		setUpExpenseRoutes(api.Group("/expenses"))
	}

	router.Run(":" + port) // Listening on the specified port
}

// getServerPort retrieves the port environment variable; defaults to "8080" if not set.
func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// setUpIncomeRoutes setups the routing for income-related API endpoints.
func setUpIncomeRoutes(income *gin.RouterGroup) {
	income.POST("/", createIncomeHandler)
	income.GET("/", getAllIncomesHandler)
	income.GET("/:id", getIncomeByIdHandler)
	income.PUT("/:id", updateIncomeByIdHandler)
	income.DELETE("/:id", deleteIncomeByIdHandler)
}

// setUpExpenseRoutes setups the routing for expense-related API endpoints.
func setUpExpenseRoutes(expenses *gin.RouterGroup) {
	expenses.POST("/", createExpenseHandler)
	expenses.GET("/", getAllExpensesHandler)
	expenses.GET("/:id", getExpenseByIdHandler)
	expenses.PUT("/:id", updateExpenseByIdHandler)
	expenses.DELETE("/:id", deleteExpenseByIdHandler)
}

func createIncomeHandler(c *gin.Context) {
	// Implementation
}

func getAllIncomesHandler(c *gin.Context) {
	// Implementation
}

func getIncomeByIdHandler(c *gin.Context) {
	// Implementation
}

func updateIncomeByIdHandler(c *gin.Context) {
	// Implementation
}

func deleteIncomeByIdHandler(c *gin.Context) {
	// Implementation
}

func createExpenseHandler(c *gin.Context) {
	// Implementation
}

func getAllExpensesHandler(c *gin.Context) {
	// Implementation
}

func getExpenseByIdHandler(c *gin.Context) {
	// Implementation
}

func updateExpenseByIdHandler(c *gin.Context) {
	// Implementation
}

func deleteExpenseByIdHandler(c *gin.Context) {
	// Implementation
}