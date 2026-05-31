package requests

import "github.com/gin-gonic/gin"

type CreateStaffRequest struct {
	CompanyID  string `bson:"company_id" json:"company_id"`
	BranchID   string `bson:"branch_id" json:"branch_id"`
	StaffName  string `bson:"staff_name" json:"staff_name"`
	Role       string `bson:"role" json:"role"`
	Department string `bson:"department" json:"department"`
	Mobile     string `bson:"mobile" json:"mobile"`
	Shift      string `bson:"shift" json:"shift"`
	Status     string `bson:"status" json:"status"`
}

type UpdateStaffRequest struct {
	StaffName  *string `bson:"staff_name,omitempty" json:"staff_name"`
	Role       *string `bson:"role,omitempty" json:"role"`
	Department *string `bson:"department,omitempty" json:"department"`
	Mobile     *string `bson:"mobile,omitempty" json:"mobile"`
	Shift      *string `bson:"shift,omitempty" json:"shift"`
	Status     *string `bson:"status,omitempty" json:"status"`
}

func NewCreateStaffRequest() *CreateStaffRequest {
	return &CreateStaffRequest{}
}

func NewUpdateStaffRequest() *UpdateStaffRequest {
	return &UpdateStaffRequest{}
}

func (r *CreateStaffRequest) Validate(c *gin.Context) error{
	return c.ShouldBind(r)
}

func(r *UpdateStaffRequest) Validate (c *gin.Context) error{
	return c.ShouldBind(r)
}