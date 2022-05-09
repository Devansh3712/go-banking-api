package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"time"

	"github.com/Devansh3712/go-banking-api/database"
	"github.com/Devansh3712/go-banking-api/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	user := models.User{Timestamp: time.Now()}
	if err = json.Unmarshal(data, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incomplete or invalid fields, check the request body.",
		})
		return
	}
	if _, err = mail.ParseAddress(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email address.",
		})
		return
	}
	accNumber, err := database.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":        fmt.Sprintf("User '%s' created.", user.Email),
		"account_number": accNumber,
	})
}

func GetUserData(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Authorization error.",
		})
		return
	}
	result, err := database.GetUserData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
