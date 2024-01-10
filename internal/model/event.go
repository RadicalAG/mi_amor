package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var EventCollectionName string = "events"

type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	StartDate   time.Time          `bson:"startDate,omitempty"`
	EndDate     time.Time          `bson:"endDate,omitempty"`
}
