package requests

import (
	"github.com/gin-gonic/gin"
)

type CreateDoctorRequest struct {
	CompanyID      string                `json:"company_id" form:"company_id" binding:"required"`
	BranchID       string                `json:"branch_id" form:"branch_id" binding:"required"`
	DoctorName     string                `json:"doctor_name" form:"doctor_name" binding:"required"`
	Specialization string                `json:"specialization" form:"specialization" binding:"required"`
	Email          string                `json:"email" form:"email"`
	Phone          string                `json:"phone" form:"phone" binding:"required"`
	OpdFees        float64               `json:"opd_fees" form:"opd_fees"`
	Status         string                `json:"status" form:"status"`
	
}

type UpdateDoctorRequest struct {
	DoctorName     *string               `json:"doctor_name,omitempty" form:"doctor_name"`
	Specialization *string               `json:"specialization,omitempty" form:"specialization"`
	Email          *string               `json:"email,omitempty" form:"email"`
	Phone          *string               `json:"phone,omitempty" form:"phone"`
	OpdFees        *float64              `json:"opd_fees,omitempty" form:"opd_fees"`
	Status         *string               `json:"status,omitempty" form:"status"`
	
}

func NewCreateDoctorRequest() *CreateDoctorRequest {
	return &CreateDoctorRequest{}
}

func NewUpdateDoctorRequest() *UpdateDoctorRequest {
	return &UpdateDoctorRequest{}
}

func (r *CreateDoctorRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}

func (r *UpdateDoctorRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}