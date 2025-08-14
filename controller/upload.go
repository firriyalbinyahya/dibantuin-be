package controller

import (
	"dibantuin-be/service"
	"dibantuin-be/utils/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadController struct {
	UploadService *service.UploadService
}

func NewUploadController(uploadService *service.UploadService) *UploadController {
	return &UploadController{UploadService: uploadService}
}

func (uc *UploadController) UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		response.BuildErrorResponse(c, errors.New("photo file is required"))
		return
	}

	filePath, err := uc.UploadService.SavePhoto(file)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		response.BuildErrorResponse(c, errors.New("failed to save and upload photo"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload success",
		"path":    filePath,
	})
}

func (uc *UploadController) UploadDocument(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.BuildErrorResponse(c, errors.New("doc file is required"))
		return
	}

	filePath, err := uc.UploadService.SaveDocument(file)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		response.BuildErrorResponse(c, errors.New("failed to save and upload doc"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload success",
		"path":    filePath,
	})
}
