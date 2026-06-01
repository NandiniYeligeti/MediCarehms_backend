package requests

import "github.com/gin-gonic/gin"

type CreateWardRequest struct {
	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	WardName  string  `bson:"ward_name" json:"ward_name"`
	Type      string  `bson:"type" json:"type"`
	TotalBeds int     `bson:"total_beds" json:"total_beds"`
	FeePerDay float64 `bson:"fee_per_day" json:"fee_per_day"`
}

type UpdateWardRequest struct {
	WardName  *string  `bson:"ward_name,omitempty" json:"ward_name,omitempty"`
	Type      *string  `bson:"type,omitempty" json:"type"`
	TotalBeds *int     `bson:"total_beds,omitempty" json:"total_beds"`
	FeePerDay *float64 `bson:"fee_per_day,omitempty" json:"fee_per_day"`
}

func NewCreateWardRequest() *CreateWardRequest {
	return &CreateWardRequest{}
}

func NewUpdateWardRequest() *UpdateWardRequest {
	return &UpdateWardRequest{}
}

func (r *CreateWardRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}

func (r *UpdateWardRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}
