package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Devansh3712/go-banking-api/database"
	"github.com/Devansh3712/go-banking-api/models"
	"github.com/gin-gonic/gin"
)

func GetUserAccountData(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Authorization error.",
		})
		return
	}
	acc, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	txn, err := database.GetTransactions(acc.Number, 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	data := map[string]interface{}{
		"account":      acc,
		"transactions": txn,
	}
	c.IndentedJSON(http.StatusOK, data)
}

func Deposit(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Authorization error.",
		})
		return
	}
	amount, got := c.GetQuery("amount")
	if !got {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query parameter amount required.",
		})
		return
	}
	parsedAmount, err := strconv.ParseFloat(amount, 64)
	amountInt := uint64(parsedAmount * 100)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Amount should be a float.",
		})
		return
	}
	result, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if err = database.UpdateAccountBalance(email, result.Balance+amountInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	txnID, err := database.CreateTransaction(models.Deposit, amount, result.Number)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to deposit amount.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":          fmt.Sprintf("Amount %s deposited to account %s.", amount, result.Number),
		"available_amount": (result.Balance + amountInt) / 100,
		"txn_id":           *txnID,
		"timestamp":        time.Now(),
	})
}

func Withdraw(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Authorization error.",
		})
		return
	}
	amount, got := c.GetQuery("amount")
	if !got {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query parameter amount required.",
		})
		return
	}
	parsedAmount, err := strconv.ParseFloat(amount, 64)
	amountInt := uint64(parsedAmount * 100)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Amount should be a float.",
		})
		return
	}
	result, err := database.GetAccountData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if result.Balance < amountInt {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":          "Withdrawal amount more than account balance.",
			"available_amount": result.Balance,
		})
		return
	}
	if err = database.UpdateAccountBalance(email, result.Balance-amountInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	txnID, err := database.CreateTransaction(models.Withdraw, amount, result.Number)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to withdraw amount.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":          fmt.Sprintf("Amount %s withdrawed from account %s.", amount, result.Number),
		"available_amount": (result.Balance - amountInt) / 100,
		"txn_id":           *txnID,
		"timestamp":        time.Now(),
	})
}
