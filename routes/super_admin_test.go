package routes

import (
	"testing"

	"github.com/NandiniYeligeti/MediCarehms_backend/models"
)

func TestBuildCompanyAdminUser(t *testing.T) {
	req := CreateCompanyRequest{
		CompanyCode: "HEALTHX",
		CompanyName: "HealthX Hospital",
		AdminEmail:  "admin@healthx.com",
		AdminPassword: "SecurePass123",
	}

	user := buildCompanyAdminUser(req, "hashed-password")

	if user.Email != req.AdminEmail {
		t.Fatalf("expected admin email %q, got %q", req.AdminEmail, user.Email)
	}
	if user.Role != "management" {
		t.Fatalf("expected management role, got %q", user.Role)
	}
	if user.CompanyCode != req.CompanyCode {
		t.Fatalf("expected company code %q, got %q", req.CompanyCode, user.CompanyCode)
	}
	if user.PasswordHash != "hashed-password" {
		t.Fatalf("expected hashed password to be preserved")
	}
	if user.Username == "" {
		t.Fatal("expected username to be populated")
	}
	if user.CompanyName != req.CompanyName {
		t.Fatalf("expected company name %q, got %q", req.CompanyName, user.CompanyName)
	}
	if _, ok := any(user).(*models.User); ok {
		// nothing to do; the helper should return a concrete *models.User
	}
}
