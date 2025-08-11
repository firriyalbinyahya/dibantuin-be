package routes

import (
	"dibantuin-be/controller"
	"dibantuin-be/middleware"
	"dibantuin-be/repository"
	"dibantuin-be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userLogRepository := repository.NewUserLogRepository(db)
	userLogService := service.NewUserLogService(userLogRepository)
	authService := service.NewAuthService(userRepository, userLogService)
	authController := controller.NewAuthController(authService)

	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.POST("/create-admin", middleware.CreateAdmin(), authController.CreateAdmin)
		auth.POST("/refresh", authController.RefreshToken)
	}
}
