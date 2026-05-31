package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"github.com/NandiniYeligeti/MediCarehms_backend/services"
	"github.com/gin-gonic/gin"
)

func UpdateDoctor(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	if companyCode == "" || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code and id are required",
		})
		return
	}

	req := requests.NewUpdateDoctorRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	service := services.NewDoctorService()

	result, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ========= staff====
func UpdateStaff(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	if companyCode == "" || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code and id are required",
		})
		return
	}

	req := requests.NewUpdateStaffRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	service := services.NewHospitalStaffService()

	result, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
