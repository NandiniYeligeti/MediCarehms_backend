package models

import (
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HospitalStaff struct {
	ID primitive.ObjectID `bson: "_id,omitempty" json: "id"`
	EntityID string `bson:"entity_id" json:"entity_id"`
	StaffCode string `bson:"staff_code" json:"staff_code"`

	CompanyID string `bson:"company_id" json:"company_if"`
	BranchID string `bson:"branch_id" json:"branch_id"`

	StaffName string `bson:"staff_name" json:"staff_name"`
	Role string `bson:"role" json:"role"`
	Department string `bson:"department" json:"department"`
	Mobile string `bson:"mobile" json:"mobile"`
	Shift string `bson:"shift" json:"shift"`
	Status         string  `bson:"status" json:"status"`

	IsDefault bool `bson:"is_default" json:"is_default"`
	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateHospitalStaff struct {
	StaffName string `bson:"staff_name" json:"staff_name"`
	Role string `bson:"role" json:"role"`
	Department string `bson:"department" json:"department"`
	Mobile string `bson:"mobile" json:"mobile"`
	Shift string `bson:"shift" json:"shift"`
	Status         string  `bson:"status" json:"status"`
}

func NewHospitalStaff() *HospitalStaff {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &HospitalStaff{
		ID:         id,
		EntityID:   id.Hex(),
		StaffCode: "DOC" + id.Hex()[18:24],
		Status:     "Active",
		IsDefault:  false,
		IsDeleted:  false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (s *HospitalStaff) Bind(req *requests.CreateStaffRequest){
	s.CompanyID = req.CompanyID
	s.BranchID = req.BranchID
	s.StaffName = req.StaffName
	s.Role = req.Role
	s.Department = req.Department
	s.Mobile = req.Mobile
	s.Shift = req.Shift
	s.Status = req.Status
	
}