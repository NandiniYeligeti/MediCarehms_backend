package routes

import (
	"context"
	"net/http"
	"strings"

	"github.com/NandiniYeligeti/MediCarehms_backend/middleware"
	"github.com/NandiniYeligeti/MediCarehms_backend/models"
	"github.com/NandiniYeligeti/MediCarehms_backend/services"
	"github.com/NandiniYeligeti/MediCarehms_backend/utils"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email/username and password are required"})
		return
	}

	identifier := strings.TrimSpace(req.Email)
	if identifier == "" {
		identifier = strings.TrimSpace(req.Username)
	}
	if identifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or username required"})
		return
	}

	// Find user by email (or username)
	ctx := context.Background()
	user, err := services.GetUserByIdentifier(ctx, identifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query user"})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Validate password
	if err := utils.CheckPassword(user.PasswordHash, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Generate JWT using user's role and company
	token, err := middleware.GenerateJWT(user.Email, user.Role, user.CompanyCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// Do not expose password hash in response
	respUser := models.User{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Role:        user.Role,
		CompanyCode: user.CompanyCode,
		CompanyName: user.CompanyName,
		Menus:       user.Menus,
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("access_token", token, 86400, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"user": respUser, "token": token})
}
