package seed

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/models"
	"github.com/NandiniYeligeti/MediCarehms_backend/services"
	"github.com/NandiniYeligeti/MediCarehms_backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EnsureSuperAdmin checks environment variables and creates a super_admin user if none exists.
func EnsureSuperAdmin() {
	email := os.Getenv("SUPER_ADMIN_EMAIL")
	password := os.Getenv("SUPER_ADMIN_PASSWORD")
	if email == "" || password == "" {
		log.Println("SUPER_ADMIN_EMAIL or SUPER_ADMIN_PASSWORD not set; skipping super admin seed")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existing, err := services.GetUserByEmail(ctx, email)
	if err != nil {
		log.Println("seed: failed to check existing super admin:", err)
		return
	}
	if existing != nil {
		log.Println("seed: super admin already exists")
		return
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		log.Println("seed: failed to hash super admin password:", err)
		return
	}

	u := &models.User{
		ID:           primitive.NewObjectID(),
		Username:     email,
		Email:        email,
		PasswordHash: hashed,
		Role:         "super_admin",
		CompanyCode:  "GLOBAL",
		CompanyName:  "MediCare Global",
	}

	if _, err := services.CreateUser(ctx, u); err != nil {
		log.Println("seed: failed to create super admin:", err)
		return
	}

	log.Println("seed: created super admin user")
}
