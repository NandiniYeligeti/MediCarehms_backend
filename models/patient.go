package models

import (
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID    string             `bson:"entity_id" json:"entity_id"`
	PatientCode string             `bson:"patient_code" json:"patient_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	PatientName string `bson:"patient_name" json:"patient_name"`
	Gender      string `bson:"gender" json:"gender"`
	Age         int    `bson:"age" json:"age"`
	Mobile      string `bson:"mobile" json:"mobile"`
	DOB         string `bson:"dob" json:"dob"`
	Address     string `bson:"address" json:"address"`
	BloodGroup  string `bson:"blood_group" json:"blood_group"`
	IsDefault   bool   `bson:"is_default" json:"is_default"`
	IsDeleted   bool   `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdatePatient struct {
	PatientName *string `bson:"patient_name,omitempty" json:"patient_name,omitempty"`
	Gender      *string `bson:"gender,omitempty" json:"gender,omitempty"`
	Age         *int    `bson:"age,omitempty" json:"age,omitempty"`
	Mobile      *string `bson:"mobile,omitempty" json:"mobile,omitempty"`
	DOB         *string `bson:"dob,omitempty" json:"dob,omitempty"`
	Address     *string `bson:"address,omitempty" json:"address,omitempty"`
	BloodGroup  *string `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
}

func NewPatient() *Patient {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &Patient{
		ID:          id,
		EntityID:    id.Hex(),
		PatientCode: "PT" + id.Hex()[18:24],
		IsDefault:   false,
		IsDeleted:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (d *Patient) Bind(req *requests.CreatePatientRequest) {
	d.CompanyID = req.CompanyID
	d.BranchID = req.BranchID
	d.PatientName = req.PatientName
	d.Gender = req.Gender
	d.Age = req.Age
	d.Mobile = req.Mobile
	d.Address = req.Address
	d.BloodGroup = req.BloodGroup
	d.DOB = req.DOB
}
