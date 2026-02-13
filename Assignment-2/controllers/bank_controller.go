package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/models"
	"gorm.io/gorm"
)

func CreateBank(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		//create empty bank object
		var b models.Bank

		//read json from request body and bind to object
		if err := c.ShouldBindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//save bacnk to db

		if err := db.Create(&b).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//send success response
		c.JSON(http.StatusCreated, b)
	}
}

func ListBanks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		//create slice to hold banks
		var banks []models.Bank

		//query db for all banks
		if err := db.Preload("Branches").Find(&banks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//send response
		c.JSON(http.StatusOK, banks)
	}
}

func UpdateBank(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		var payload models.Bank
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var bank models.Bank
		if err := db.First(&bank, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "bank not found"})
			return
		}
		bank.Name = payload.Name
		bank.Code = payload.Code
		if err := db.Save(&bank).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bank)
	}
}

func DeleteBank(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		if err := db.Delete(&models.Bank{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "bank deleted"})
	}
}
