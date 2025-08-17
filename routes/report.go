package routes

import (
	"dibantuin-be/controller"
	"dibantuin-be/middleware"
	"dibantuin-be/repository"
	"dibantuin-be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupReportRoutes(router *gin.RouterGroup, db *gorm.DB) {
	donationProgramRepository := repository.NewDonationProgramRepository(db)
	donationTransactionRepository := repository.NewDonationTransactionRepository(db)
	reportRepository := repository.NewReportRepository(db)
	reportService := service.NewReportService(donationProgramRepository, donationTransactionRepository, reportRepository)
	reportController := controller.NewReportController(reportService)

	report := router.Group("/report")
	{
		report.GET("/global", middleware.AdminOnly(), reportController.GetGlobalReport)
		report.GET("/program/:id", middleware.AuthUserMiddleware(), reportController.GetDonationProgramReport)
	}
}
