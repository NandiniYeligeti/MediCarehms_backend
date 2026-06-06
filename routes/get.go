package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/services"
	"github.com/gin-gonic/gin"
)

func GetDoctors(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	service := services.NewDoctorService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetDoctorByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	service := services.NewDoctorService()

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetDoctorByEntityID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	entityID := c.Param("entity_id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "entity_id is required",
		})
		return
	}

	service := services.NewDoctorService()

	data, err := service.GetByEntityID(ctx, companyCode, entityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// ========= staff =======

func GetStaffs(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	service := services.NewHospitalStaffService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetStaffByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	service := services.NewHospitalStaffService()

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetStaffByEntityID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	entityID := c.Param("entity_id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "entity_id is required",
		})
		return
	}

	service := services.NewHospitalStaffService()

	data, err := service.GetByEntityID(ctx, companyCode, entityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// =========ward====
func GetWards(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	service := services.NewWardService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetWardByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	service := services.NewWardService()

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetWardByEntityID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	entityID := c.Param("entity_id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "entity_id is required",
		})
		return
	}

	service := services.NewWardService()

	data, err := service.GetByEntityID(ctx, companyCode, entityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// ========patient====
func GetPatients(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("Company Code:", c.Param("company_code"))

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required"})
		return
	}

	service := services.NewPatientService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetPatientByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	service := services.NewPatientService()

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetPatientByEntityID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	entityID := c.Param("entity_id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "entity_id is required",
		})
		return
	}

	service := services.NewPatientService()

	data, err := service.GetByEntityID(ctx, companyCode, entityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// =====birth certificate=====
func GetBirthCertificates(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	service := services.NewBirthCertificateService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetBirthCertificateByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	service := services.NewBirthCertificateService()

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetBirthCertificateByEntityID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	entityID := c.Param("entity_id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "entity_id is required",
		})
		return
	}

	service := services.NewBirthCertificateService()

	data, err := service.GetByEntityID(ctx, companyCode, entityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// opd booking
func GetOPDBookings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	service := services.NewOPDBookingService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetOPDBookingByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	service := services.NewOPDBookingService()

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetOPDBookingByEntityID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	entityID := c.Param("entity_id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "entity_id is required",
		})
		return
	}

	service := services.NewOPDBookingService()

	data, err := service.GetByEntityID(ctx, companyCode, entityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
