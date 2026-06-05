package requests

import (
	

	"github.com/gin-gonic/gin"
)

type CreateBirthCertificateRequest struct {
	CompanyID               string  `json:"company_id" form:"company_id" binding:"required"`
	BranchID                string  `json:"branch_id" form:"branch_id" binding:"required"`
	BabyName                string  `json:"baby_name" form:"baby_name" binding:"required"`
	Gender                  string  `json:"gender" form:"gender" binding:"required"`
	DOB                     string  `json:"dob" form:"dob" binding:"required"`
	Time                    string  `json:"time" form:"time" binding:"required"`
	Weight                  float64 `json:"weight" form:"weight" binding:"required"`
	MotherName              string  `json:"mother_name" form:"mother_name" binding:"required"`
	FatherName              string  `json:"father_name" form:"father_name" binding:"required"`
	DoctorDirectoryEntityID string  `json:"doctor_directory_entity_id" form:"doctor_directory_entity_id" binding:"required"`
}

type UpdateBirthCertificateRequest struct {
	BabyName                *string  `json:"baby_name,omitempty" form:"baby_name,omitempty"`
	Gender                  *string  `json:"gender,omitempty" form:"gender,omitempty"`
	DOB                     *string  `json:"dob,omitempty" form:"dob,omitempty"`
	Time                    *string  `json:"time,omitempty" form:"time,omitempty"`
	Weight                  *float64 `json:"weight,omitempty" form:"weight,omitempty"`
	MotherName              *string  `json:"mother_name,omitempty" form:"mother_name,omitempty"`
	FatherName              *string  `json:"father_name,omitempty" form:"father_name,omitempty"`
	DoctorDirectoryEntityID *string  `json:"doctor_directory_entity_id,omitempty" form:"doctor_directory_entity_id,omitempty"`
}

func NewCreateBirthCertificateRequest() *CreateBirthCertificateRequest {
	return &CreateBirthCertificateRequest{}
}

func NewUpdateBirthCertificateRequest() *UpdateBirthCertificateRequest {
	return &UpdateBirthCertificateRequest{}
}

func (r *CreateBirthCertificateRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}

func (r *UpdateBirthCertificateRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}
