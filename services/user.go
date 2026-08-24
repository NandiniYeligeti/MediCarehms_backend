package services

import (
	"context"
	"errors"
	"time"

	"github.com/NandiniYeligeti/MediCarehms_backend/models"
	"github.com/NandiniYeligeti/MediCarehms_backend/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollection = "users"

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := storage.GetCollection(UsersCollection)
	var u models.User
	err := coll.FindOne(cctx, bson.M{"email": email}).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

// GetUserByIdentifier tries to find a user by email first, then by username.
func GetUserByIdentifier(ctx context.Context, identifier string) (*models.User, error) {
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := storage.GetCollection(UsersCollection)
	var u models.User
	// try email
	err := coll.FindOne(cctx, bson.M{"email": identifier}).Decode(&u)
	if err == nil {
		return &u, nil
	}
	if err != mongo.ErrNoDocuments {
		return nil, err
	}
	// try username
	err = coll.FindOne(cctx, bson.M{"username": identifier}).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func CreateUser(ctx context.Context, user *models.User) (*primitive.ObjectID, error) {
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := storage.GetCollection(UsersCollection)
	res, err := coll.InsertOne(cctx, user)
	if err != nil {
		return nil, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("invalid inserted id")
	}
	return &oid, nil
}

func UpdateUserPassword(ctx context.Context, userID primitive.ObjectID, hashed string) error {
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := storage.GetCollection(UsersCollection)
	_, err := coll.UpdateOne(cctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"password_hash": hashed}})
	return err
}

func GetUsersByCompany(ctx context.Context, companyCode string) ([]models.User, error) {
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := storage.GetCollection(UsersCollection)
	cursor, err := coll.Find(cctx, bson.M{"company_code": companyCode})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(cctx)

	var out []models.User
	for cursor.Next(cctx) {
		var u models.User
		if err := cursor.Decode(&u); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, nil
}
