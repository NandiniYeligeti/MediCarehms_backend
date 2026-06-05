package models

import (
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BirthCertificate struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID string             `bson:"entity_id" json:"entity_id"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	BabyName                string    `bson:"baby_name" json:"baby_name"`
	Gender                  string    `bson:"gender" json:"gender"`
	DOB                     string    `bson:"dob" json:"dob"`
	Time                    string    `bson:"time" json:"time"`
	Weight                  float64   `bson:"weight" json:"weight"`
	MotherName              string    `bson:"mother_name" json:"mother_name"`
	FatherName              string    `bson:"father_name" json:"father_name"`
	DoctorDirectoryEntityID string    `bson:"doctor_directory_entity_id" json:"doctor_directory_entity_id"`

	IsDefault bool `bson:"is_default" json:"is_default"`
	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateBirthCertificate struct {
	BabyName                *string    `bson:"baby_name,omitempty" json:"baby_name,omitempty"`
	Gender                  *string    `bson:"gender,omitempty" json:"gender,omitempty"`
	DOB                     *string    `bson:"dob,omitempty" json:"dob,omitempty"`
	Time                    *string    `bson:"time,omitempty" json:"time,omitempty"`
	Weight                  *float64   `bson:"weight,omitempty" json:"weight,omitempty"`
	MotherName              *string    `bson:"mother_name,omitempty" json:"mother_name,omitempty"`
	FatherName              *string    `bson:"father_name,omitempty" json:"father_name,omitempty"`
	DoctorDirectoryEntityID *string    `bson:"doctor_directory_entity_id,omitempty" json:"doctor_directory_entity_id,omitempty"`
}

func NewBirthCertificate() *BirthCertificate {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &BirthCertificate{
		ID:        id,
		EntityID:  id.Hex(),
		IsDefault: false,
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (d *BirthCertificate) Bind(req *requests.CreateBirthCertificateRequest) {
	d.CompanyID = req.CompanyID
	d.BranchID = req.BranchID
	d.BabyName = req.BabyName
	d.Gender = req.Gender
	d.DOB = req.DOB
	d.Time = req.Time
	d.Weight = req.Weight
	d.MotherName = req.MotherName
	d.FatherName = req.FatherName
	d.DoctorDirectoryEntityID = req.DoctorDirectoryEntityID
}
