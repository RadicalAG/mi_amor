package repository

import (
	"context"
	"log"
	"radical/red_letter/internal/internal_error"
	"radical/red_letter/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewUserRepository(client *mongo.Client, databaseName, collectionName string) *userRepository {
	return &userRepository{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
}

func (r *userRepository) getCollection() *mongo.Collection {
	return r.client.Database(r.databaseName).Collection(r.collectionName)
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	collection := r.getCollection()

	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return internal_error.InternalServerError("error creating user")
	}

	return nil
}
