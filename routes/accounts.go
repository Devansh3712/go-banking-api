package routes

import (
	"net/http"

	"github.com/Devansh3712/go-banking-api/database"
	"github.com/gin-gonic/gin"
)

func GetUserAccountData(c *gin.Context) {
	email := c.MustGet("email").(string)
	result, err := database.GetUserAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
