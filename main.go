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
	err := database.Migrate()
	if err != nil {
		panic(err)
	}
	fmt.Println("Migrations successfull.")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/api", root)
	router.GET("/api/login", routes.AuthHandler)
	router.POST("/api/users", routes.CreateUser)
	router.GET("/api/users", middleware.JWTAuthMiddleware(), routes.GetUserData)
	router.GET("/api/accounts", middleware.JWTAuthMiddleware(), routes.GetUserAccountData)

	fmt.Println("API running on http://localhost:8080")
	router.Run("localhost:8080")
}
