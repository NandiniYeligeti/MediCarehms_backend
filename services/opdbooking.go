package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/models"
	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"github.com/NandiniYeligeti/MediCarehms_backend/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const OPDBookingCollection = "opdbookings"

// ======Service interface=======
type OPDBookingService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateOPDBookingRequest) (*models.OPDBooking, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.OPDBooking, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.OPDBooking, error)
	GetByEntityID(ctx context.Context, companyCode string, entityID string) (*models.OPDBooking, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateOPDBookingRequest) (*models.OPDBooking, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ===Service struct =====
type opdBookingService struct{}

func NewOPDBookingService() OPDBookingService {
	return &opdBookingService{}
}

// ========== create appointment=========
func (s *opdBookingService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateOPDBookingRequest,
) (*models.OPDBooking, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(OPDBookingCollection)

	// Generate Token Number
	count, err := collection.CountDocuments(
		ctx,
		bson.M{
			"doctor_entity_id": req.DoctorEntityID,
			"appointment_date": req.AppointmentDate,
			"is_deleted":       bson.M{"$ne": true},
		},
	)

	if err != nil {
		return nil, err
	}

	booking := models.NewOPDBooking()
	booking.Bind(req)

	// Auto Token
	booking.TokenNumber = int(count) + 1

	_, err = collection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

// ======= get all appointment =======
func (s *opdBookingService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.OPDBooking, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(OPDBookingCollection)

	opts := options.Find().SetSort(
		bson.M{"created_at": -1},
	)

	cursor, err := collection.Find(
		ctx,
		bson.M{
			"is_deleted": bson.M{"$ne": true},
		},
		opts,
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var bookings []*models.OPDBooking

	if err := cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}

// ========== get appointment by id ======
func (s *opdBookingService) GetByID(ctx context.Context, companyCode string, id string) (*models.OPDBooking, error) {
	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(OPDBookingCollection)

	filter := bson.M{
		"entity_id":  id,
		"is_deleted": bson.M{"$ne": true},
	}

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{
			"_id":        oid,
			"is_deleted": bson.M{"$ne": true},
		}
	}

	var booking models.OPDBooking

	err := collection.FindOne(ctx, filter).Decode(&booking)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("booking not found")
	}
	return &booking, err
}

// ======get appointment by entity id=========
func (s *opdBookingService) GetByEntityID(ctx context.Context, companyCode string, entityID string) (*models.OPDBooking, error) {
	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(OPDBookingCollection)

	filter := bson.M{
		"entity_id":  entityID,
		"is_deleted": bson.M{"$ne": true},
	}

	var booking models.OPDBooking

	err := collection.FindOne(ctx, filter).Decode(&booking)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("booking not found")
	}
	return &booking, err
}

// ===== update appointment =====
func (s *opdBookingService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateOPDBookingRequest,
) (*models.OPDBooking, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(OPDBookingCollection)

	updateFields := bson.M{}

	if req.PatientEntityID != nil {
		updateFields["patient_entity_id"] = *req.PatientEntityID
	}

	if req.DoctorEntityID != nil {
		updateFields["doctor_entity_id"] = *req.DoctorEntityID
	}

	if req.AppointmentDate != nil {
		updateFields["appointment_date"] = *req.AppointmentDate
	}

	if req.TimeSlot != nil {
		updateFields["time_slot"] = *req.TimeSlot
	}

	if req.Status != nil {
		updateFields["status"] = *req.Status
	}

	updateFields["updated_at"] = time.Now().UTC()

	filter := bson.M{
		"entity_id": id,
	}

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{
			"_id": oid,
		}
	}

	result, err := collection.UpdateOne(
		ctx,
		filter,
		bson.M{
			"$set": updateFields,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("appointment not found")
	}

	var updated models.OPDBooking

	_ = collection.FindOne(
		ctx,
		filter,
	).Decode(&updated)

	return &updated, nil
}

// delete appointment
func (s *opdBookingService) Delete(ctx context.Context, companyCode string, id string) error {
	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(OPDBookingCollection)

	filter := bson.M{
		"entity_id": id,
	}

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{
			"_id": oid,
		}
	}

	update := bson.M{
		"$set": bson.M{
			"is_deleted": true,
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := collection.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("appointment not found")
	}
	return nil
}
