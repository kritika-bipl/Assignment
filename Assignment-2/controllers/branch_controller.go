package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/models"
	"gorm.io/gorm"
)

func CreateBranch(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b models.Branch

		if err := c.ShouldBindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&b).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, b)

	}
}

func ListBranchesByBank(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bankId := c.Param("bankId")
		var branches []models.Branch
		if err := db.Where("bank_id = ?", bankId).Find(&branches).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, branches)
	}
}
