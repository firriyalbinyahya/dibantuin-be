package routes

import (
	"dibantuin-be/controller"
	"dibantuin-be/middleware"
	"dibantuin-be/repository"
	"dibantuin-be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupDonationProgramRoutes(router *gin.RouterGroup, db *gorm.DB) {
	donationProgramRepository := repository.NewDonationProgramRepository(db)
	userLogRepository := repository.NewUserLogRepository(db)
	verificationRepository := repository.NewVerificationProgramRepository(db)
	userLogService := service.NewUserLogService(userLogRepository)
	donationProgramService := service.NewDonationProgramService(db, donationProgramRepository, verificationRepository, userLogService)
	donationProgramController := controller.NewDonationProgramController(donationProgramService)

	auth := router.Group("/program")
	{
		auth.POST("/request", middleware.AuthUserMiddleware(), donationProgramController.RequestProgram)
		auth.PATCH("/:id/verify", middleware.AdminOnly(), donationProgramController.VerifyProgram)
	}
}
