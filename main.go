package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Devansh3712/go-banking-api/database"
	"github.com/Devansh3712/go-banking-api/middleware"
	"github.com/Devansh3712/go-banking-api/routes"
	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Banking API is working.",
		"timestamp": time.Now(),
	})
}

func main() {
	if err := database.MigratePostgres(); err != nil {
		panic(err)
	}
	if err := database.MigrateImmuDB(); err != nil {
		panic(err)
	}
	fmt.Println("Migrations successfull.")

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	v1 := app.Group("/api/v1")
	{
		v1.GET("/", root)
		v1.GET("/login", routes.AuthHandler)
		user := v1.Group("/user")
		{
			user.GET("/", middleware.JWTAuthMiddleware(), routes.GetUserData)
			user.POST("/", routes.CreateUser)
		}
		account := v1.Group("/account")
		{
			account.GET("/", middleware.JWTAuthMiddleware(), routes.GetUserAccountData)
			account.GET("/deposit", middleware.JWTAuthMiddleware(), routes.Deposit)
			account.GET("/withdraw", middleware.JWTAuthMiddleware(), routes.Withdraw)
		}
		transaction := v1.Group("/transaction")
		{
			transaction.GET("/", middleware.JWTAuthMiddleware(), routes.GetTransactions)
			transaction.GET("/deposit", middleware.JWTAuthMiddleware(), routes.GetDeposits)
			transaction.GET("/withdraw", middleware.JWTAuthMiddleware(), routes.GetWithdawals)
		}
	}

	fmt.Println("API running on http://localhost:8000")
	app.Run("localhost:8000")
}
