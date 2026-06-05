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

const PatientCollection = "patient"

// ================== PHONE VALIDATION ==================

var patientPhoneRegex = regexp.MustCompile(`^\d{10}$`)

func isValidPatientPhone(mobile string) bool {
	return patientPhoneRegex.MatchString(mobile)
}

// ================== SERVICE INTERFACE ==================

type PatientService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreatePatientRequest) (*models.Patient, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.Patient, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.Patient, error)
	GetByEntityID(ctx context.Context, companyCode string, entityID string) (*models.Patient, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdatePatientRequest) (*models.Patient, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================== SERVICE STRUCT ==================

type patientService struct{}

func NewPatientService() PatientService {
	return &patientService{}
}

// ================== CREATE PATIENT ==================

func (s *patientService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreatePatientRequest,
) (*models.Patient, error) {

	if !isValidPatientPhone(req.Mobile) {
		return nil, errors.New("mobile number must be exactly 10 digits")
	}

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(PatientCollection)

	if count, _ := collection.CountDocuments(ctx, bson.M{
		"mobile": req.Mobile,
		"is_deleted": bson.M{"$ne": true},
	}); count > 0 {
		return nil, errors.New("patient mobile already exists")
	}

	patient := models.NewPatient()
	patient.Bind(req)

	_, err := collection.InsertOne(ctx, patient)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

// ================== GET ALL PATIENTS ==================

func (s *patientService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.Patient, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(PatientCollection)

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

	var patients []*models.Patient

	if err := cursor.All(ctx, &patients); err != nil {
		return nil, err
	}

	return patients, nil
}

// ================== GET PATIENT BY ID ==================

func (s *patientService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.Patient, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(PatientCollection)

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

	var patient models.Patient

	err := collection.FindOne(ctx, filter).Decode(&patient)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("patient not found")
	}

	return &patient, err
}

// ================== GET PATIENT BY ENTITY ID ==================

func (s *patientService) GetByEntityID(
	ctx context.Context,
	companyCode string,
	entityID string,
) (*models.Patient, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(PatientCollection)

	filter := bson.M{
		"entity_id": entityID,
		"is_deleted": bson.M{"$ne": true},
	}

	var patient models.Patient

	err := collection.FindOne(ctx, filter).Decode(&patient)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("patient not found")
	}

	return &patient, err
}

// ================== UPDATE PATIENT ==================

func (s *patientService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdatePatientRequest,
) (*models.Patient, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(PatientCollection)

	updateFields := bson.M{}

	if req.PatientName != nil {
		updateFields["patient_name"] = *req.PatientName
	}

	if req.Gender != nil {
		updateFields["gender"] = *req.Gender
	}

	if req.Age != nil {
		updateFields["age"] = *req.Age
	}

	if req.Mobile != nil {
		if !isValidPatientPhone(*req.Mobile) {
			return nil, errors.New("mobile number must be exactly 10 digits")
		}
		updateFields["mobile"] = *req.Mobile
	}

	if req.Address != nil {
		updateFields["address"] = *req.Address
	}

	if req.BloodGroup != nil {
		updateFields["blood_group"] = *req.BloodGroup
	}

	if req.DOB != nil {
		dob, err := time.Parse("2006-01-02", *req.DOB)
		if err != nil {
			return nil, errors.New("invalid dob format, use YYYY-MM-DD")
		}
		updateFields["dob"] = dob
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
		bson.M{"$set": updateFields},
	)

	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("patient not found")
	}

	var updated models.Patient

	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

// ================== DELETE PATIENT ==================

func (s *patientService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(PatientCollection)

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
		return errors.New("patient not found")
	}

	return nil
}