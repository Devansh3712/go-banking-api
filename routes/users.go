package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Devansh3712/go-banking-api/database"
	"github.com/Devansh3712/go-banking-api/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid parameter: %s", err),
		})
		return
	}
	user := models.User{Timestamp: time.Now()}
	err = json.Unmarshal(data, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err := database.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("%s user created.", user.Email),
	})
}

func GetUserData(c *gin.Context) {
	email := c.MustGet("email").(string)
	result, err := database.GetUserData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
