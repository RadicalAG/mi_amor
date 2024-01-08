package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserCollectionName string = "user"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}
