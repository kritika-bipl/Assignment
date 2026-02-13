package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/services"
	"gorm.io/gorm"
)

type txRequest struct {
	AccountID uint    `json:"account_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required,gt=0"`
}

func DepositHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r txRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := services.Deposit(db, r.AccountID, r.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deposit successful"})
	}
}

func WithdrawHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r txRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := services.Withdraw(db, r.AccountID, r.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "withdraw successful"})
	}
}

func ListTransactions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		_, trs, err := services.GetAccountWithTransactions(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, trs)
	}
}
