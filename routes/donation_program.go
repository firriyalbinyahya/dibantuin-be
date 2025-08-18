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

	program := router.Group("/program")
	{
		program.GET("/", middleware.AuthUserMiddleware(), donationProgramController.ListDonationPrograms)
		program.GET("/:id/admin", middleware.AdminOnly(), donationProgramController.GetDonationProgramDetail)
		program.GET("/:id", middleware.AuthUserMiddleware(), donationProgramController.GetDonationProgramDetailForUser)
		program.POST("/request", middleware.AuthUserMiddleware(), donationProgramController.RequestProgram)
		program.DELETE("/:id", middleware.AuthUserMiddleware(), donationProgramController.DeleteDonationProgram)
		program.PATCH("/update/:id", middleware.AuthUserMiddleware(), donationProgramController.UpdateDonationProgram)
		program.PATCH("/:id/verify", middleware.AdminOnly(), donationProgramController.VerifyProgram)
	}
}
