package repository

import (
	"context"
	"fmt"
	"log"
	"radical/red_letter/internal/internal_error"
	"radical/red_letter/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type eventRepository struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewEventRepository(client *mongo.Client, databaseName, collectionName string) EventRepository {
	return &eventRepository{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

type EventRepository interface {
	CreateEvent(ctx context.Context, event *model.Event) (string, error)
	GetEventByID(ctx context.Context, eventID string) (*model.Event, error)
}

func (r *eventRepository) getCollection() *mongo.Collection {
	return r.client.Database(r.databaseName).Collection(r.collectionName)
}
func (r *eventRepository) CreateEvent(ctx context.Context, event *model.Event) (string, error) {
	collection := r.getCollection()

	// Set a new ObjectID if not provided
	if event.ID.IsZero() {
		event.ID = primitive.NewObjectID()
	}

	result, err := collection.InsertOne(context.TODO(), event)
	if err != nil {
		log.Printf("Error creating event: %v\n", err)
		return "", internal_error.InternalServerError("error creating event")
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Printf("failed to convert InsertedID to ObjectID")
		return "", internal_error.InternalServerError("error creating event")
	}

	return insertedID.Hex(), nil
}

func (r *eventRepository) GetEventByID(ctx context.Context, eventID string) (*model.Event, error) {
	collection := r.getCollection()

	objectID, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		log.Printf("invalid ObjectID format")
		return nil, internal_error.NotFoundError("event")
	}

	filter := bson.M{"_id": objectID}

	var event model.Event
	err = collection.FindOne(ctx, filter).Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("event not found")
			return nil, internal_error.NotFoundError("event")
		}
		return nil, internal_error.InternalServerError("")
	}

	return &event, nil
}

func (r *eventRepository) UpdateEvent(ctx context.Context, eventID string, updatedEvent *model.Event) error {
	collection := r.getCollection()

	objectID, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return fmt.Errorf("invalid ObjectID format")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedEvent}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("event not updated")
	}

	return nil
}

func (r *eventRepository) DeleteEvent(ctx context.Context, eventID string) error {
	collection := r.getCollection()

	objectID, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return fmt.Errorf("invalid ObjectID format")
	}

	filter := bson.M{"_id": objectID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("event not deleted")
	}

	return nil
}
