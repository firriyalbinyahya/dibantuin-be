package routes

import (
	"dibantuin-be/controller"
	"dibantuin-be/middleware"
	"dibantuin-be/repository"
	"dibantuin-be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCategoryRoutes(router *gin.RouterGroup, db *gorm.DB) {
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	category := router.Group("/categories")
	{
		category.POST("", middleware.AdminOnly(), categoryController.CreateCategory)
		category.GET("", middleware.AdminOnly(), categoryController.GetCategories)
		category.PUT("/:id", middleware.AdminOnly(), categoryController.UpdateCategory)
		category.DELETE("/:id", middleware.AdminOnly(), categoryController.DeleteCategory)
	}
}
