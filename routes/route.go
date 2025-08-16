package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api/v1")
	{
		SetupAuthRoutes(api, db)
		SetupDonationProgramRoutes(api, db)
		SetupDonationTransactionRoutes(api, db)
		SetupUserLogRoutes(api, db)
		SetupCategoryRoutes(api, db)
		SetupUploadRoutes(api)
	}
}
