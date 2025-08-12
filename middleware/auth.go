package middleware

import (
	"dibantuin-be/config"
	"dibantuin-be/utils/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func APIKeyChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" {
			c.Set("api-key", false)
			c.Next()
			return
		}

		if apiKey != config.Config.AdminAPIKey {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Invalid API key",
			})
			c.Abort()
			return
		}

		c.Set("api-key", true)
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		user, err := auth.VerifyToken(tokenParts[1], false)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		if user.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Set("user_id", user.ID)
		c.Next()
	}
}

func CreateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cek API Key dulu
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey != "" {
			if apiKey == config.Config.AdminAPIKey {
				c.Set("api-key", true)
				c.Next()
				return
			} else {
				c.JSON(http.StatusForbidden, gin.H{"error": "Invalid API key"})
				c.Abort()
				return
			}
		}

		// Kalau API Key kosong maka cek JWT
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		user, err := auth.VerifyToken(tokenParts[1], false)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		if user.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Set("role", "admin")
		c.Set("api-key", false)
		c.Next()
	}
}

func AuthUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		user, err := auth.VerifyToken(tokenParts[1], false)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Set("user_id", user.ID)
		c.Next()
	}
}
