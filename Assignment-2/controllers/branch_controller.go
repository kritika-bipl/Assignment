package controllers

import (
	"net/http"
	"strconv"

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

func UpdateBranch(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		var payload models.Branch
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var br models.Branch
		if err := db.First(&br, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "branch not found"})
			return
		}
		br.Name = payload.Name
		br.IFSCCode = payload.IFSCCode
		br.Address = payload.Address
		if payload.BankID != 0 {
			br.BankID = payload.BankID
		}
		if err := db.Save(&br).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, br)
	}
}

func DeleteBranch(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		if err := db.Delete(&models.Branch{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "branch deleted"})
	}
}
