package routes

import (
	"dibantuin-be/controller"
	"dibantuin-be/middleware"
	"dibantuin-be/repository"
	"dibantuin-be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupDonationTransactionRoutes(router *gin.RouterGroup, db *gorm.DB) {
	donationTransactionRepository := repository.NewDonationTransactionRepository(db)
	userLogRepository := repository.NewUserLogRepository(db)
	transactionVerificationRepository := repository.NewTransactionVerificationRepository(db)
	userLogService := service.NewUserLogService(userLogRepository)
	donationTransactionService := service.NewDonationTransactionService(donationTransactionRepository, transactionVerificationRepository, userLogService)
	donationTransactionController := controller.NewDonationTransactionController(donationTransactionService)

	program := router.Group("/donation")
	{
		program.POST("/", middleware.AuthUserMiddleware(), donationTransactionController.CreateMoneyDonationTransaction)
		program.PATCH("/:id/verify", middleware.AdminOnly(), donationTransactionController.VerifyTransactionDonation)
		program.GET("/", middleware.AdminOnly(), donationTransactionController.ListDonationTransactions)
		program.GET("/user-history", middleware.AuthUserMiddleware(), donationTransactionController.ListDonationTransactionsUser)
	}
}
