package routes

import (
	"dibantuin-be/controller"
	"dibantuin-be/middleware"
	"dibantuin-be/repository"
	"dibantuin-be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserLogRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userLogRepository := repository.NewUserLogRepository(db)
	userLogService := service.NewUserLogService(userLogRepository)
	userLogController := controller.NewUserLogController(userLogService)

	userLog := router.Group("/user-logs")
	{
		userLog.GET("/:id", middleware.AdminOnly(), userLogController.GetUserLogs)
	}
}
