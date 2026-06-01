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

const WardCollection = "wards"

// ================== SERVICE INTERFACE ==================

type WardService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateWardRequest) (*models.WardSetup, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.WardSetup, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.WardSetup, error)
	GetByEntityID(ctx context.Context, companyCode string, entityID string) (*models.WardSetup, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateWardRequest) (*models.WardSetup, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================== SERVICE STRUCT ==================

type wardService struct{}

func NewWardService() WardService {
	return &wardService{}
}

// ================== CREATE WARD ==================

func (s *wardService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateWardRequest,
) (*models.WardSetup, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(WardCollection)

	if count, _ := collection.CountDocuments(ctx, bson.M{
		"ward_name":  req.WardName,
		"is_deleted": bson.M{"$ne": true},
	}); count > 0 {
		return nil, errors.New("ward name already exists")
	}

	ward := models.NewWardSetup()
	ward.Bind(req)

	_, err := collection.InsertOne(ctx, ward)
	if err != nil {
		return nil, err
	}

	return ward, nil
}

// ================== GET ALL WARDS ==================

func (s *wardService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.WardSetup, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(WardCollection)

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

	var wards []*models.WardSetup

	if err := cursor.All(ctx, &wards); err != nil {
		return nil, err
	}

	return wards, nil
}

// ================== GET WARD BY ID ==================

func (s *wardService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.WardSetup, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(WardCollection)

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

	var ward models.WardSetup

	err := collection.FindOne(ctx, filter).Decode(&ward)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("ward not found")
	}

	return &ward, err
}

// ================== GET WARD BY ENTITY ID ==================

func (s *wardService) GetByEntityID(
	ctx context.Context,
	companyCode string,
	entityID string,
) (*models.WardSetup, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(WardCollection)

	filter := bson.M{
		"entity_id":  entityID,
		"is_deleted": bson.M{"$ne": true},
	}

	var ward models.WardSetup

	err := collection.FindOne(ctx, filter).Decode(&ward)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("ward not found")
	}

	return &ward, err
}

// ================== UPDATE WARD ==================

func (s *wardService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateWardRequest,
) (*models.WardSetup, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(WardCollection)

	updateFields := bson.M{}

	if req.WardName != nil {
		updateFields["ward_name"] = *req.WardName
	}

	if req.Type != nil {
		updateFields["type"] = *req.Type
	}

	if req.TotalBeds != nil {
		updateFields["total_beds"] = *req.TotalBeds
	}

	if req.FeePerDay != nil {
		updateFields["fee_per_day"] = *req.FeePerDay
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
		return nil, errors.New("ward not found")
	}

	var updated models.WardSetup
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

// ================== DELETE WARD ==================

func (s *wardService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(WardCollection)

	filter := bson.M{"entity_id": id}

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}

	result, err := collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": bson.M{
			"is_deleted": true,
			"updated_at": time.Now().UTC(),
		}},
	)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("ward not found")
	}

	return nil
}
