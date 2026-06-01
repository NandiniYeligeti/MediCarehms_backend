package models

import (
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WardSetup struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID string             `bson:"entity_id" json:"entity_id"`
	//	WardCode string             `bson:"ward_code" json:"ward_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	WardName  string  `bson:"ward_name" json:"ward_name"`
	Type      string  `bson:"type" json:"type"`
	TotalBeds int     `bson:"total_beds" json:"total_beds"`
	FeePerDay float64 `bson:"fee_per_day" json:"fee_per_day"`

	IsDefault bool `bson:"is_default" json:"is_default"`
	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateWardSetup struct {
	WardName  *string  `bson:"ward_name,omitempty" json:"ward_name,omitempty"`
	Type      *string  `bson:"type,omitempty" json:"type"`
	TotalBeds *int     `bson:"total_beds,omitempty" json:"total_beds"`
	FeePerDay *float64 `bson:"fee_per_day,omitempty" json:"fee_per_day"`
}

func NewWardSetup() *WardSetup {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &WardSetup{
		ID:        id,
		EntityID:  id.Hex(),
		IsDefault: false,
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (w *WardSetup) Bind(req *requests.CreateWardRequest) {
	w.CompanyID = req.CompanyID
	w.BranchID = req.BranchID
	w.WardName = req.WardName
	w.Type = req.Type
	w.TotalBeds = req.TotalBeds
	w.FeePerDay = req.FeePerDay
}
