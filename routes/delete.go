package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/services"
	"github.com/gin-gonic/gin"
)

func DeleteDoctor(c *gin.Context) {
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

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Doctor deleted successfully",
	})
}

// ======staff=======
func DeleteHospitalStaff(c *gin.Context) {
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

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hospital staff deleted successfully",
	})
}

// =====ward==========
func DeleteWard(c *gin.Context) {
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

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ward deleted successfully",
	})
}

//======== Patient===========
func DeletePatient (c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required"})
		return
	}

	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" :"id is required"		})
			return
	}

	service := services.NewPatientService()

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Patient deleted successfully",
	})
}
//======birth certificate ======
func DeleteBirthCertificate(c *gin.Context) {
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
	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Birth certificate deleted successfully",
	})
}