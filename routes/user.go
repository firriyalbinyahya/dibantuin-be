package routes

import (
	"dibantuin-be/controller"
	"dibantuin-be/middleware"
	"dibantuin-be/repository"
	"dibantuin-be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	user := router.Group("/user")
	{
		user.GET("/:id", middleware.AuthUserMiddleware(), userController.GetUserByID)
		user.PUT("/:id", middleware.AuthUserMiddleware(), userController.UpdateUser)
		user.DELETE("/:id", middleware.AuthUserMiddleware(), userController.DeleteUser)
	}
}
