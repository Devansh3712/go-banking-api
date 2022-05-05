package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Devansh3712/go-banking-api/auth"
	"github.com/Devansh3712/go-banking-api/database"
	"github.com/Devansh3712/go-banking-api/models"
	"github.com/gin-gonic/gin"
)

// Authorize user credentials to return a JWT and set
// JWT as a cookie.
func AuthHandler(c *gin.Context) {
	var user models.UserAuth
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &user)
	err = database.AuthUser(&user)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": fmt.Sprintf("Authentication failed: %s", err),
		})
		return
	}
	tokenString, _ := auth.GenerateToken(user.Email)
	cookie := http.Cookie{
		Name:    "JWT",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour),
	}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, gin.H{
		"message": "User authorized.",
		"token":   tokenString,
	})
}
