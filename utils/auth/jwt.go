package auth

import (
	"dibantuin-be/config"
	"dibantuin-be/entity"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.User, isRefresh bool) (string, *time.Time, error) {
	var (
		secret string
		expiry int
	)

	if isRefresh {
		secret = config.Config.JWT.RefreshSecret
		expiry = config.Config.JWT.RefreshExpiryInSec
	} else {
		secret = config.Config.JWT.AccessSecret
		expiry = config.Config.JWT.AccessExpiryInSec
	}

	expiredAt := time.Now().Add(time.Second * time.Duration(expiry))
	claims := &JWTClaim{
		ID:    fmt.Sprint(user.ID),
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "firriyal-bin-yahya",
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, &expiredAt, nil
}

func VerifyToken(tokenString string, isRefresh bool) (*entity.User, error) {
	var secret string
	if isRefresh {
		secret = config.Config.JWT.RefreshSecret
	} else {
		secret = config.Config.JWT.AccessSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		user := &entity.User{}
		idInt, _ := strconv.ParseUint(claims.ID, 10, 64)
		user.ID = idInt
		user.Email = claims.Email
		user.Role = claims.Role

		return user, nil
	}

	return nil, err
}
