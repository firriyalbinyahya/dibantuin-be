package controller

import (
	"dibantuin-be/entity"
	"dibantuin-be/service"
	"dibantuin-be/utils/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{UserService: service}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, errors.New("invalid user ID"))
		return
	}

	user, err := uc.UserService.GetUserByID(userID)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "User retrieved successfully", user, nil)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, errors.New("invalid user ID"))
		return
	}

	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		response.BuildErrorResponse(c, fmt.Errorf("user not authenticated"))
		return
	}

	if userIDFromToken != userID {
		response.BuildErrorResponse(c, errors.New("you are not authorized to update this user"))
		return
	}

	var request entity.UserUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, errors.New("invalid request body"))
		return
	}

	err = uc.UserService.UpdateUser(userID, &request)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "User updated successfully", nil, nil)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	targetUserIDStr := c.Param("id")
	targetUserID, err := strconv.ParseUint(targetUserIDStr, 10, 64)
	if err != nil {
		response.BuildErrorResponse(c, errors.New("invalid user id"))
		return
	}

	actingUserID, _ := c.Get("user_id")
	actingUserRole, _ := c.Get("user_role")

	err = uc.UserService.DeleteUser(actingUserID.(uint64), targetUserID, actingUserRole.(string))
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "User deleted successfully", nil, nil)
}
