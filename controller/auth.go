package controller

import (
	"dibantuin-be/entity"
	"dibantuin-be/service"
	"dibantuin-be/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (ac *AuthController) Register(c *gin.Context) {
	var req entity.Register

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	err := ac.Service.Register(&req)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "registration successful", nil, nil)
}

func (ac *AuthController) CreateAdmin(c *gin.Context) {
	var req entity.Register

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	fromAPIKeyVal, exists := c.Get("api-key")
	if !exists {
		fromAPIKeyVal = false
	}
	fromAPIKey, ok := fromAPIKeyVal.(bool)
	if !ok {
		fromAPIKey = false
	}

	roleVal, _ := c.Get("role")
	role, _ := roleVal.(string)
	var idAdminInt uint64 = 0

	userRaw, _ := c.Get("currentUser")

	user, ok := userRaw.(*entity.User)
	if ok {
		idAdminInt = user.ID
	}

	err := ac.Service.CreateaAdmin(&req, fromAPIKey, role, idAdminInt)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "The admin account was created successfully", nil, nil)

}

func (ac *AuthController) Login(c *gin.Context) {
	var request entity.Login

	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	result, err := ac.Service.Login(&request)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Login successful", result, nil)

}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	var req entity.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	resp, err := ac.Service.RefreshToken(req.RefreshToken)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Refresh token successful", resp, nil)
}

func (ac *AuthController) Logout(c *gin.Context) {
	userID := c.GetUint64("user_id")

	err := ac.Service.Logout(userID)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Logout successful", nil, nil)
}
