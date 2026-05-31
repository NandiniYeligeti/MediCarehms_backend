package services

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/models"
	"github.com/NandiniYeligeti/MediCarehms_backend/requests"
	"github.com/NandiniYeligeti/MediCarehms_backend/storage"
    

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DoctorCollection = "doctors"

// ================== PHONE VALIDATION ==================

var doctorPhoneRegex = regexp.MustCompile(`^\d{10}$`)

func isValidDoctorPhone(phone string) bool {
	return doctorPhoneRegex.MatchString(phone)
}

// ================== SERVICE INTERFACE ==================

type DoctorService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateDoctorRequest) (*models.DoctorDirectory, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.DoctorDirectory, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.DoctorDirectory, error)
	GetByEntityID(ctx context.Context, companyCode string, entityID string) (*models.DoctorDirectory, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateDoctorRequest) (*models.DoctorDirectory, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================== SERVICE STRUCT ==================

type doctorService struct{}

func NewDoctorService() DoctorService {
	return &doctorService{}
}

// ================== CREATE DOCTOR ==================

func (s *doctorService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateDoctorRequest,
) (*models.DoctorDirectory, error) {

	if !isValidDoctorPhone(req.Phone) {
		return nil, errors.New("phone number must be exactly 10 digits")
	}

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(DoctorCollection)

	if count, _ := collection.CountDocuments(ctx, bson.M{
		"phone": req.Phone,
		"is_deleted": bson.M{"$ne": true},
	}); count > 0 {
		return nil, errors.New("doctor phone already exists")
	}

	doctor := models.NewDoctorDirectory()
	doctor.Bind(req)

	_, err := collection.InsertOne(ctx, doctor)
	if err != nil {
		return nil, err
	}

	return doctor, nil
}

// ================== GET ALL DOCTORS ==================

func (s *doctorService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.DoctorDirectory, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(DoctorCollection)

	opts := options.Find().SetSort(bson.M{
		"created_at": -1,
	})

	cursor, err := collection.Find(
		ctx,
		bson.M{"is_deleted": bson.M{"$ne": true}},
		opts,
	)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var doctors []*models.DoctorDirectory

	if err := cursor.All(ctx, &doctors); err != nil {
		return nil, err
	}

	return doctors, nil
}

// ================== GET DOCTOR BY ID ==================

func (s *doctorService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.DoctorDirectory, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(DoctorCollection)

	filter := bson.M{
		"entity_id": id,
		"is_deleted": bson.M{"$ne": true},
	}

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{
			"_id": oid,
			"is_deleted": bson.M{"$ne": true},
		}
	}

	var doctor models.DoctorDirectory

	err := collection.FindOne(ctx, filter).Decode(&doctor)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("doctor not found")
	}

	return &doctor, err
}

// ================== GET DOCTOR BY ENTITY ID ==================

func (s *doctorService) GetByEntityID(
	ctx context.Context,
	companyCode string,
	entityID string,
) (*models.DoctorDirectory, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(DoctorCollection)

	filter := bson.M{
		"entity_id": entityID,
		"is_deleted": bson.M{"$ne": true},
	}

	var doctor models.DoctorDirectory

	err := collection.FindOne(ctx, filter).Decode(&doctor)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("doctor not found")
	}

	return &doctor, err
}

// ================== UPDATE DOCTOR ==================

func (s *doctorService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateDoctorRequest,
) (*models.DoctorDirectory, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(DoctorCollection)

	updateFields := bson.M{}

	if req.DoctorName != nil {
		updateFields["doctor_name"] = *req.DoctorName
	}

	if req.Specialization != nil {
		updateFields["specialization"] = *req.Specialization
	}


	if req.Email != nil {
		updateFields["email"] = *req.Email
	}

	if req.Phone != nil {
		if !isValidDoctorPhone(*req.Phone) {
			return nil, errors.New("phone number must be exactly 10 digits")
		}
		updateFields["phone"] = *req.Phone
	}

	if req.OpdFees != nil {
		updateFields["opd_fees"] = *req.OpdFees
	}

	if req.Status != nil {
		updateFields["status"] = *req.Status
	}

	updateFields["updated_at"] = time.Now().UTC()

	filter := bson.M{"entity_id": id}

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}

	result, err := collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": updateFields},
	)

	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("doctor not found")
	}

	var updated models.DoctorDirectory
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

// ================== DELETE DOCTOR ==================

func (s *doctorService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(DoctorCollection)

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
		return errors.New("doctor not found")
	}

	return nil
}