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

const BirthCertificateCollection = "birthcertificate"

// ================= SERVICE INTERFACE =================

type BirthCertificateService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateBirthCertificateRequest) (*models.BirthCertificate, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.BirthCertificate, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.BirthCertificate, error)
	GetByEntityID(ctx context.Context, companyCode string, entityID string) (*models.BirthCertificate, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateBirthCertificateRequest) (*models.BirthCertificate, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================= SERVICE STRUCT =================

type birthCertificateService struct{}

func NewBirthCertificateService() BirthCertificateService {
	return &birthCertificateService{}
}

// ================= CREATE =================

func (s *birthCertificateService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateBirthCertificateRequest,
) (*models.BirthCertificate, error) {

	db := storage.GetMongo()

	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(BirthCertificateCollection)

	certificate := models.NewBirthCertificate()
	certificate.Bind(req)

	_, err := collection.InsertOne(ctx, certificate)
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

// ================= GET ALL =================

func (s *birthCertificateService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.BirthCertificate, error) {

	db := storage.GetMongo()

	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(BirthCertificateCollection)

	opts := options.Find().SetSort(
		bson.M{"created_at": -1},
	)

	cursor, err := collection.Find(
		ctx,
		bson.M{"is_deleted": bson.M{"$ne": true}},
		opts,
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var certificates []*models.BirthCertificate

	if err := cursor.All(ctx, &certificates); err != nil {
		return nil, err
	}

	return certificates, nil
}

// ================= GET BY ID =================

func (s *birthCertificateService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.BirthCertificate, error) {

	db := storage.GetMongo()

	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(BirthCertificateCollection)

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

	var certificate models.BirthCertificate

	err := collection.FindOne(ctx, filter).Decode(&certificate)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("birth certificate not found")
	}

	return &certificate, err
}

// ================= GET BY ENTITY ID =================

func (s *birthCertificateService) GetByEntityID(
	ctx context.Context,
	companyCode string,
	entityID string,
) (*models.BirthCertificate, error) {

	db := storage.GetMongo()

	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(BirthCertificateCollection)

	filter := bson.M{
		"entity_id":  entityID,
		"is_deleted": bson.M{"$ne": true},
	}

	var certificate models.BirthCertificate

	err := collection.FindOne(ctx, filter).Decode(&certificate)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("birth certificate not found")
	}

	return &certificate, err
}

// ================= UPDATE =================

func (s *birthCertificateService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateBirthCertificateRequest,
) (*models.BirthCertificate, error) {

	db := storage.GetMongo()

	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(BirthCertificateCollection)

	updateFields := bson.M{}

	if req.BabyName != nil {
		updateFields["baby_name"] = *req.BabyName
	}

	if req.Gender != nil {
		updateFields["gender"] = *req.Gender
	}

	if req.DOB != nil {
		updateFields["dob"] = *req.DOB
	}

	if req.Time != nil {
		updateFields["time"] = *req.Time
	}

	if req.Weight != nil {
		updateFields["weight"] = *req.Weight
	}

	if req.MotherName != nil {
		updateFields["mother_name"] = *req.MotherName
	}

	if req.FatherName != nil {
		updateFields["father_name"] = *req.FatherName
	}

	if req.DoctorDirectoryEntityID != nil {
		updateFields["doctor_directory_entity_id"] = *req.DoctorDirectoryEntityID
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
		return nil, errors.New("birth certificate not found")
	}

	var updated models.BirthCertificate

	_ = collection.FindOne(
		ctx,
		filter,
	).Decode(&updated)

	return &updated, nil
}

// ================= DELETE =================

func (s *birthCertificateService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()

	collection := db.Database(
		fmt.Sprintf("company_%s", companyCode),
	).Collection(BirthCertificateCollection)

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
		return errors.New("birth certificate not found")
	}

	return nil
}
