package models

import (
	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DoctorDirectory struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID string             `bson:"entity_id" json:"entity_id"`
	DoctorCode string           `bson:"doctor_code" json:"doctor_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	DoctorName     string  `bson:"doctor_name" json:"doctor_name"`
	Specialization string  `bson:"specialization" json:"specialization"`
	
	Email          string  `bson:"email" json:"email"`
	Phone          string  `bson:"phone" json:"phone"`
	OpdFees        float64 `bson:"opd_fees" json:"opd_fees"`
	Status         string  `bson:"status" json:"status"`


	IsDefault bool `bson:"is_default" json:"is_default"`
	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateDoctorDirectory struct {
	DoctorName     *string  `bson:"doctor_name,omitempty" json:"doctor_name,omitempty"`
	Specialization *string  `bson:"specialization,omitempty" json:"specialization,omitempty"`
	
	Email          *string  `bson:"email,omitempty" json:"email,omitempty"`
	Phone          *string  `bson:"phone,omitempty" json:"phone,omitempty"`
	OpdFees        *float64 `bson:"opd_fees,omitempty" json:"opd_fees,omitempty"`
	Status         *string  `bson:"status,omitempty" json:"status,omitempty"`
}

func NewDoctorDirectory() *DoctorDirectory {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &DoctorDirectory{
		ID:         id,
		EntityID:   id.Hex(),
		DoctorCode: "DOC" + id.Hex()[18:24],
		Status:     "Active",
		IsDefault:  false,
		IsDeleted:  false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (d *DoctorDirectory) Bind(req *requests.CreateDoctorRequest) {
	d.CompanyID = req.CompanyID
	d.BranchID = req.BranchID
	d.DoctorName = req.DoctorName
	d.Specialization = req.Specialization
	d.Email = req.Email
	d.Phone = req.Phone
	d.OpdFees = req.OpdFees
	d.Status = req.Status
}
