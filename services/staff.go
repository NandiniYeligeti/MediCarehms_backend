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

const HospitalStaffCollection = "hospital_staff"

// ================== PHONE VALIDATION ==================

var staffMobileRegex = regexp.MustCompile(`^\d{10}$`)

func isValidStaffMobile(mobile string) bool {
	return staffMobileRegex.MatchString(mobile)
}

// ================== SERVICE INTERFACE ==================

type HospitalStaffService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateStaffRequest) (*models.HospitalStaff, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.HospitalStaff, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.HospitalStaff, error)
	GetByEntityID(ctx context.Context, companyCode string, entityID string) (*models.HospitalStaff, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateStaffRequest) (*models.HospitalStaff, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================== SERVICE STRUCT ==================

type hospitalStaffService struct{}

func NewHospitalStaffService() HospitalStaffService {
	return &hospitalStaffService{}
}

// ================== CREATE STAFF ==================

func (s *hospitalStaffService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateStaffRequest,
) (*models.HospitalStaff, error) {

	if !isValidStaffMobile(req.Mobile) {
		return nil, errors.New("mobile number must be exactly 10 digits")
	}

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode)).Collection(HospitalStaffCollection)

	if count, _ := collection.CountDocuments(ctx, bson.M{
		"mobile":     req.Mobile,
		"is_deleted": bson.M{"$ne": true},
	}); count > 0 {
		return nil, errors.New("staff mobile already exists")
	}

	staff := models.NewHospitalStaff()
	staff.Bind(req)

	_, err := collection.InsertOne(ctx, staff)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

// ================== GET ALL STAFF ==================

func (s *hospitalStaffService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.HospitalStaff, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(HospitalStaffCollection)

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

	var staffs []*models.HospitalStaff

	if err := cursor.All(ctx, &staffs); err != nil {
		return nil, err
	}

	return staffs, nil
}

// ================== GET STAFF BY ID ==================

func (s *hospitalStaffService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.HospitalStaff, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(HospitalStaffCollection)

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

	var staff models.HospitalStaff

	err := collection.FindOne(ctx, filter).Decode(&staff)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("staff not found")
	}

	return &staff, err
}

// ================== GET STAFF BY ENTITY ID ==================

func (s *hospitalStaffService) GetByEntityID(
	ctx context.Context,
	companyCode string,
	entityID string,
) (*models.HospitalStaff, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(HospitalStaffCollection)

	filter := bson.M{
		"entity_id":  entityID,
		"is_deleted": bson.M{"$ne": true},
	}

	var staff models.HospitalStaff

	err := collection.FindOne(ctx, filter).Decode(&staff)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("staff not found")
	}

	return &staff, err
}

// ================== UPDATE STAFF ==================

func (s *hospitalStaffService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateStaffRequest,
) (*models.HospitalStaff, error) {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(HospitalStaffCollection)

	updateFields := bson.M{}

	if req.StaffName != nil {
		updateFields["staff_name"] = *req.StaffName
	}

	if req.Role != nil {
		updateFields["role"] = *req.Role
	}

	if req.Department != nil {
		updateFields["department"] = *req.Department
	}

	if req.Mobile != nil {
		if !isValidStaffMobile(*req.Mobile) {
			return nil, errors.New("mobile number must be exactly 10 digits")
		}
		updateFields["mobile"] = *req.Mobile
	}

	if req.Shift != nil {
		updateFields["shift"] = *req.Shift
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
		bson.M{"$set": updateFields},
	)

	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("staff not found")
	}

	var updated models.HospitalStaff

	_ = collection.FindOne(
		ctx,
		filter,
	).Decode(&updated)

	return &updated, nil
}

// ================== DELETE STAFF ==================

func (s *hospitalStaffService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(HospitalStaffCollection)

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
		return errors.New("staff not found")
	}

	return nil
}
