package entity

import "time"

type Register struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Role           string     `json:"role"`
	AccessToken    string     `json:"access_token"`
	AccessExpired  *time.Time `json:"access_expired"`
	RefreshToken   string     `json:"refresh_token"`
	RefreshExpired *time.Time `json:"refresh_expired"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
