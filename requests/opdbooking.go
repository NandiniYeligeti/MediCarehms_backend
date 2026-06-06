package requests

import "github.com/gin-gonic/gin"

type CreateOPDBookingRequest struct {
	CompanyID       string `bson:"company_id" json:"company_id" binding:"required"`
	BranchID        string `bson:"branch_id" json:"branch_id" binding:"required"`
	PatientEntityID string `bson:"patient_entity_id" json:"patient_entity_id" binding:"required"`
	DoctorEntityID  string `bson:"doctor_entity_id" json:"doctor_entity_id" binding:"required"`
	AppointmentDate string `bson:"appointment_date" json:"appointment_date" binding:"required"`
	TimeSlot        string `bson:"time_slot" json:"time_slot" binding:"required"`
	Status          string `bson:"status" json:"status" binding:"required"`
}

type UpdateOPDBookingRequest struct {
	PatientEntityID *string `bson:"patient_entity_id" json:"patient_entity_id"`
	DoctorEntityID  *string `bson:"doctor_entity_id" json:"doctor_entity_id"`
	AppointmentDate *string `bson:"appointment_date" json:"appointment_date"`
	TimeSlot        *string `bson:"time_slot" json:"time_slot"`
	Status          *string `bson:"status" json:"status"`
}

func NewCreateOPDBookingRequest() *CreateOPDBookingRequest {
	return &CreateOPDBookingRequest{}
}

func NewUpdateOPDBookingRequest() *UpdateOPDBookingRequest {
	return &UpdateOPDBookingRequest{}
}

func (r *CreateOPDBookingRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}

func (r *UpdateOPDBookingRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}
