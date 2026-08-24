package routes

import (
	"context"
	"net/http"
	"strings"

	"github.com/NandiniYeligeti/MediCarehms_backend/middleware"
	"github.com/NandiniYeligeti/MediCarehms_backend/models"
	"github.com/NandiniYeligeti/MediCarehms_backend/services"
	"github.com/NandiniYeligeti/MediCarehms_backend/storage"
	"github.com/NandiniYeligeti/MediCarehms_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Company struct {
	ID          string `json:"id"`
	CompanyCode string `json:"company_code"`
	CompanyName string `json:"company_name"`
	AdminID     string `json:"admin_id,omitempty"`
	AdminEmail  string `json:"admin_email,omitempty"`
}

type CreateCompanyRequest struct {
	CompanyCode   string `json:"company_code" binding:"required"`
	CompanyName   string `json:"company_name" binding:"required"`
	AdminEmail    string `json:"admin_email" binding:"required,email"`
	AdminPassword string `json:"admin_password" binding:"required,min=6"`
}

var demoCompanies = []Company{
	{ID: uuid.NewString(), CompanyCode: "DEFAULT_COMPANY", CompanyName: "MediCare HMS", AdminID: "admin", AdminEmail: "admin@medicare.com"},
	{ID: uuid.NewString(), CompanyCode: "GLOBAL", CompanyName: "MediCare Global", AdminID: "superadmin", AdminEmail: "superadmin@medicare.com"},
}

func buildCompanyAdminUser(req CreateCompanyRequest, passwordHash string) *models.User {
	username := strings.ToLower(strings.ReplaceAll(req.CompanyCode, " ", "_"))
	if username == "" {
		username = strings.ToLower(strings.ReplaceAll(req.CompanyName, " ", "_"))
	}
	if username == "" {
		username = "administrator"
	}

	return &models.User{
		Username:     username,
		Email:        strings.TrimSpace(req.AdminEmail),
		PasswordHash: passwordHash,
		Role:         "management",
		CompanyCode:  strings.TrimSpace(req.CompanyCode),
		CompanyName:  strings.TrimSpace(req.CompanyName),
	}
}

func getCompaniesFromStore(ctx context.Context) ([]Company, error) {
	coll := storage.GetCollection("companies")
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var companies []Company
	for cursor.Next(ctx) {
		var company Company
		if err := cursor.Decode(&company); err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	if len(companies) == 0 {
		return demoCompanies, nil
	}
	return companies, nil
}

func ensureCompanyCodeAvailable(ctx context.Context, companyCode string) (bool, error) {
	coll := storage.GetCollection("companies")
	count, err := coll.CountDocuments(ctx, bson.M{"company_code": strings.TrimSpace(companyCode)})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func saveCompany(ctx context.Context, company Company) error {
	coll := storage.GetCollection("companies")
	_, err := coll.InsertOne(ctx, company)
	return err
}

func requireSuperAdmin(c *gin.Context) (*middleware.Claims, bool) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, false
	}

	claims, ok := user.(*middleware.Claims)
	if !ok || claims.Role != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: super admin only"})
		return nil, false
	}
	return claims, true
}

func GetCompanies(c *gin.Context) {
	if _, ok := requireSuperAdmin(c); !ok {
		return
	}

	ctx := context.Background()
	companies, err := getCompaniesFromStore(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch companies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"companies": companies})
}

func CreateCompany(c *gin.Context) {
	if _, ok := requireSuperAdmin(c); !ok {
		return
	}

	var req CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_code, company_name, admin_email and admin_password are required"})
		return
	}

	ctx := context.Background()
	available, err := ensureCompanyCodeAvailable(ctx, req.CompanyCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to validate company code"})
		return
	}
	if !available {
		c.JSON(http.StatusConflict, gin.H{"error": "company code already exists"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.AdminPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to secure admin password"})
		return
	}

	adminUser := buildCompanyAdminUser(req, hashedPassword)
	if _, err := services.CreateUser(ctx, adminUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create company admin account"})
		return
	}

	company := Company{
		ID:          uuid.NewString(),
		CompanyCode: strings.TrimSpace(req.CompanyCode),
		CompanyName: strings.TrimSpace(req.CompanyName),
		AdminID:     adminUser.Email,
		AdminEmail:  adminUser.Email,
	}
	if err := saveCompany(ctx, company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save company"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"company": company, "admin": gin.H{"email": adminUser.Email, "username": adminUser.Username}})
}
