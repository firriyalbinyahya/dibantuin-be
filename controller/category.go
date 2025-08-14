package controller

import (
	"dibantuin-be/entity"
	"dibantuin-be/service"
	"dibantuin-be/utils/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService *service.CategoryService
}

func NewCategoryController(categoryService *service.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var request entity.CategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	category, err := cc.CategoryService.CreateCategory(&request)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "success to create category.", category, nil)
}

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, errors.New("invalid id"))
		return
	}

	var request entity.CategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	category, err := cc.CategoryService.UpdateCategory(id, &request)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "success to update category.", category, nil)
}

func (cc *CategoryController) GetCategories(c *gin.Context) {
	categories, err := cc.CategoryService.GetCategories()
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "success to get categories.", categories, nil)
}

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, errors.New("invalid id"))
		return
	}

	if err := cc.CategoryService.DeleteCategory(id); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "success to delete category.", nil, nil)
}
