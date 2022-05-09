package routes

import (
	"net/http"
	"strconv"

	"github.com/Devansh3712/go-banking-api/database"
	"github.com/Devansh3712/go-banking-api/models"
	"github.com/gin-gonic/gin"
)

// Get all types of transactions, default limit is 10.
func GetTransactions(c *gin.Context) {
	email := c.MustGet("email").(string)
	acc, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	var TxnLimit int
	limit, got := c.GetQuery("limit")
	if !got {
		TxnLimit = 10
	} else {
		TxnLimit, err = strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
	}
	txn, err := database.GetTransactions(acc.Number, TxnLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	data := map[string]interface{}{
		"transactions": txn,
	}
	c.IndentedJSON(http.StatusOK, data)
}

func GetWithdawals(c *gin.Context) {
	email := c.MustGet("email").(string)
	acc, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	var TxnLimit int
	limit, got := c.GetQuery("limit")
	if !got {
		TxnLimit = 10
	} else {
		TxnLimit, err = strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
	}
	txn, err := database.GetTransactionsByType(models.Withdraw, acc.Number, TxnLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	data := map[string]interface{}{
		"transactions": txn,
	}
	c.IndentedJSON(http.StatusOK, data)
}

func GetDeposits(c *gin.Context) {
	email := c.MustGet("email").(string)
	acc, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	var TxnLimit int
	limit, got := c.GetQuery("limit")
	if !got {
		TxnLimit = 10
	} else {
		TxnLimit, err = strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
	}
	txn, err := database.GetTransactionsByType(models.Deposit, acc.Number, TxnLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	data := map[string]interface{}{
		"transactions": txn,
	}
	c.IndentedJSON(http.StatusOK, data)
}
