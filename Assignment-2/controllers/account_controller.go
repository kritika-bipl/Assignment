package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/models"
	"github.com/kritika-bipl/Assignment/Assignment-2/services"
	"gorm.io/gorm"
)

func OpenAccountHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req models.Account

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := services.OpenAccount(db, &req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, req)
	}
}

func GetAccountHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		acc, _, err := services.GetAccountWithTransactions(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, acc)
	}
}

func GetAccountsByCustomerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("customerId")
		id, _ := strconv.Atoi(idStr)
		accs, err := services.GetAccountsByCustomer(db, uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, accs)
	}
}
