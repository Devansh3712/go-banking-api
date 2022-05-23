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
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, models.AuthorizationError)
		return
	}
	acc, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AccountFetchError)
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
	c.JSON(http.StatusOK, data)
}

func GetTransactionByID(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, models.AuthorizationError)
		return
	}
	acc, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AccountFetchError)
		return
	}
	txnID := c.Param("txnID")
	txn, err := database.GetTransactionByID(acc.Number, txnID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, txn)
}

func GetWithdrawals(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, models.AuthorizationError)
		return
	}
	acc, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AccountFetchError)
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
	c.JSON(http.StatusOK, data)
}

func GetDeposits(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, models.AuthorizationError)
		return
	}
	acc, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AccountFetchError)
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
	c.JSON(http.StatusOK, data)
}
