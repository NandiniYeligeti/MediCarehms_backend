package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"github.com/NandiniYeligeti/MediCarehms_backend/services"
	"github.com/gin-gonic/gin"
)

func CreateDoctor(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// company code from URL
	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	// bind request
	req := requests.NewCreateDoctorRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// service call
	service := services.NewDoctorService()

	doctor, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, doctor)
}

//======== staff====

func CreateStaff(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// company code from URL
	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	// bind request
	req := requests.NewCreateStaffRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// service call
	service := services.NewHospitalStaffService()

	staff, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, staff)
}
//====ward ========
func CreateWard(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// company code from URL
	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	// bind request
	req := requests.NewCreateWardRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// service call
	service := services.NewWardService()

	ward, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ward)
}
//====patient====
func CreatePatient(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// company code from URL
	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	// bind request
	req := requests.NewCreatePatientRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// service call
	service := services.NewPatientService()

	patient, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, patient)
}
//======birth certificate ======
func CreateBirthCertificate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// company code from URL
	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	// bind request
	req := requests.NewCreateBirthCertificateRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// service call
	service := services.NewBirthCertificateService()	

	birthCertificate, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, birthCertificate)
}	