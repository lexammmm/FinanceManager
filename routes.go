package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"sync"
)

var (
	largeStructPool = sync.Pool{
		New: func() interface{} {
			return &LargeStruct{}
		},
	}
)

type LargeStruct struct {
}

func main() {
	router := gin.Default()

	port := getServerPort()

	api := router.Group("/api")
	{
		setUpIncomeRoutes(api.Group("/incomes"))
		setUpExpenseRoutes(api.Group("/expenses"))
	}

	router.Run(":" + port)
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func setUpIncomeRoutes(income *gin.RouterGroup) {
	income.POST("/", createIncomeHandler)
}

func setUpExpenseRoutes(expenses *gin.RouterGroup) {
}

func createIncomeHandler(c *gin.Context) {
	ls := largeStructPool.Get().(*LargeStruct)
	defer largeStructPool.Put(ls)
}