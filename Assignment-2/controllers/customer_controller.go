package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/models"
	"gorm.io/gorm"
)

func CreateCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var cust models.Customer

		if err := c.ShouldBindJSON(&cust) ; err != nil {
			c.JSON(http.StatusBadRequest , gin.H { "error" : err.Error()})
			return
		}

		if err := db.Create(&cust).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, cust)
	}
}

//we are using first here because uit returns only one record
//also it retirn not found error while find() return empty slice

func GetCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var cust models.Customer
		if err := db.Preload("Accounts").Preload("Loans").First(&cust, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cust)
	}
}
