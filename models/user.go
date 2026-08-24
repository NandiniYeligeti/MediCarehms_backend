package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username     string             `bson:"username,omitempty" json:"username"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash string             `bson:"password_hash" json:"-"`
	Role         string             `bson:"role" json:"role"`
	CompanyCode  string             `bson:"company_code,omitempty" json:"company_code,omitempty"`
	CompanyName  string             `bson:"company_name,omitempty" json:"company_name,omitempty"`
	Menus        []string           `bson:"menus,omitempty" json:"menus,omitempty"`
}
