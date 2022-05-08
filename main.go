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

	app.GET("/api", root)
	app.GET("/api/login", routes.AuthHandler)

	app.POST("/api/user", routes.CreateUser)
	app.GET("/api/user", middleware.JWTAuthMiddleware(), routes.GetUserData)

	app.GET("/api/account", middleware.JWTAuthMiddleware(), routes.GetUserAccountData)
	app.GET("/api/account/deposit", middleware.JWTAuthMiddleware(), routes.Deposit)
	app.GET("/api/account/withdraw", middleware.JWTAuthMiddleware(), routes.Withdraw)

	fmt.Println("API running on http://localhost:8000")
	app.Run("localhost:8000")
}
