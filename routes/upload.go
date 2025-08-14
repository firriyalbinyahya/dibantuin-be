package routes

import (
	"dibantuin-be/controller"
	"dibantuin-be/middleware"
	"dibantuin-be/service"

	"github.com/gin-gonic/gin"
)

func SetupUploadRoutes(router *gin.RouterGroup) {
	uploadService := service.NewUploadService("./uploads")
	uploadController := controller.NewUploadController(uploadService)

	upload := router.Group("/upload")
	{
		upload.POST("/photo", middleware.AuthUserMiddleware(), uploadController.UploadPhoto)
		upload.POST("/document", middleware.AuthUserMiddleware(), uploadController.UploadDocument)
	}
}
