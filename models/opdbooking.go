package models

import (
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OPDBooking struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID        string             `bson:"entity_id" json:"entity_id"`
	AppointmentCode string             `bson:"appointment_code" json:"appointment_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	PatientEntityID string `bson:"patient_entity_id" json:"patient_entity_id"`
	DoctorEntityID  string `bson:"doctor_entity_id" json:"doctor_entity_id"`
	AppointmentDate string `bson:"appointment_date" json:"appointment_date"`
	TimeSlot        string `bson:"time_slot" json:"time_slot"`
	TokenNumber     int    `bson:"token_number" json:"token_number"`
	Status          string `bson:"status" json:"status"`

	IsDefault bool `bson:"is_default" json:"is_default"`
	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateOPDBooking struct {
	PatientEntityID string `bson:"patient_entity_id" json:"patient_entity_id"`
	DoctorEntityID  string `bson:"doctor_entity_id" json:"doctor_entity_id"`
	AppointmentDate string `bson:"appointment_date" json:"appointment_date"`
	TimeSlot        string `bson:"time_slot" json:"time_slot"`
	TokenNumber     int    `bson:"token_number" json:"token_number"`
	Status          string `bson:"status" json:"status"`
}

func NewOPDBooking() *OPDBooking {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &OPDBooking{
		ID:              id,
		EntityID:        id.Hex(),
		AppointmentCode: "APT" + id.Hex()[18:24],
		Status:          "Booked",
		IsDefault:       false,
		IsDeleted:       false,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func (d *OPDBooking) Bind(req *requests.CreateOPDBookingRequest) {
	d.CompanyID = req.CompanyID
	d.BranchID = req.BranchID
	d.PatientEntityID = req.PatientEntityID
	d.DoctorEntityID = req.DoctorEntityID
	d.AppointmentDate = req.AppointmentDate
	d.TimeSlot = req.TimeSlot
	d.Status = req.Status
}
