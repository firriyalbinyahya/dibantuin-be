package controller

import (
	"dibantuin-be/service"
	"dibantuin-be/utils/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserLogController struct {
	UserLogService *service.UserLogService
}

func NewUserLogController(service *service.UserLogService) *UserLogController {
	return &UserLogController{UserLogService: service}
}

func (ulc *UserLogController) GetUserLogs(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, errors.New("invalid user id"))
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	paginationData, err := ulc.UserLogService.GetUserLogs(userID, limit, page)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Success get data of user logs", paginationData, nil)

}
