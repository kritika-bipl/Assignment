package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/models"
	"gorm.io/gorm"
)

func CreateBank(db *gorm.DB) gin.HandlerFunc {
   return func(c *gin.Context) {
	       
	     //create empty bank object
		 var b models.Bank 

		 //read json from request body and bind to object
		 if err := c.ShouldBindJSON(&b); err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
			return
		 }

		 //save bacnk to db

		 if err := db.Create(&b).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		//send success response
		c.JSON(http.StatusCreated , b)
   }
}

func ListBanks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		//create slice to hold banks
		var banks []models.Bank

		//query db for all banks
		if err := db.Find(&banks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//send response
		c.JSON(http.StatusOK, banks)
	}
}
