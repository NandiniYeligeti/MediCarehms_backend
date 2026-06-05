package requests

import (
	"github.com/gin-gonic/gin"
)

type CreatePatientRequest struct {
	CompanyID   string `json:"company_id" form:"company_id" binding:"required"`
	BranchID    string `json:"branch_id" form:"branch_id" binding:"required"`
	PatientName string `json:"patient_name" form:"patient_name" binding:"required"`
	Gender      string `json:"gender" form:"gender" binding:"required"`
	Age         int    `json:"age" form:"age" binding:"required"`
	Mobile      string `json:"mobile" form:"mobile" binding:"required"`
	DOB         string `json:"dob" form:"dob" binding:"required"`
	Address     string `json:"address" form:"address" binding:"required"`
	BloodGroup  string `json:"blood_group" form:"blood_group" binding:"required"`
}

type UpdatePatientRequest struct {
	PatientName *string `json:"patient_name,omitempty" form:"patient_name"`
	Gender      *string `json:"gender,omitempty" form:"gender"`
	Age         *int    `json:"age,omitempty" form:"age"`
	Mobile      *string `json:"mobile,omitempty" form:"mobile"`
	DOB         *string `json:"dob,omitempty" form:"dob"`
	Address     *string `json:"address,omitempty" form:"address"`
	BloodGroup  *string `json:"blood_group,omitempty" form:"blood_group"`
}

func NewCreatePatientRequest() *CreatePatientRequest {
	return &CreatePatientRequest{}
}

func NewUpdatePatientRequest() *UpdatePatientRequest {
	return &UpdatePatientRequest{}
}

func (r *CreatePatientRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}

func (r *UpdatePatientRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}
