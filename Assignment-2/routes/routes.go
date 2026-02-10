package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/controllers"
	"gorm.io/gorm"
)


func RegisterRoutes(r *gin.Engine , db *gorm.DB){

	api := r.Group("/api")
	{    
		//bank api
		api.POST("/banks", controllers.CreateBank(db))
		api.GET("/banks", controllers.ListBanks(db))

		//branches
		api.POST("/branches", controllers.CreateBranch(db))
		api.GET("/branches/:bankId", controllers.ListBranchesByBank(db))

	}
}