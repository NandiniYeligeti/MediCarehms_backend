package routes

import (
	"context"
	"net/http"

	"github.com/NandiniYeligeti/MediCarehms_backend/models"
	"github.com/NandiniYeligeti/MediCarehms_backend/services"
	"github.com/NandiniYeligeti/MediCarehms_backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
	CompanyCode string `json:"company_code"`
}

func CreateUserHandler(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashed,
		Role:         req.Role,
		CompanyCode:  req.CompanyCode,
	}

	ctx := context.Background()
	oid, err := services.CreateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": oid.Hex()})
}

func GetUsersHandler(c *gin.Context) {
	company := c.Param("company_code")
	ctx := context.Background()
	users, err := services.GetUsersByCompany(ctx, company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	// hide password hashes
	for i := range users {
		users[i].PasswordHash = ""
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func DeleteUserHandler(c *gin.Context) {
	// not implemented for brevity; would remove user by id
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func ChangePasswordHandler(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password required"})
		return
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	hashed, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash"})
		return
	}
	ctx := context.Background()
	if err := services.UpdateUserPassword(ctx, oid, hashed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
