package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Username string
	Password string
	Uri      string
}

func NewMongoDB(username, password, uri string) *MongoDB {
	return &MongoDB{
		Username: username,
		Password: password,
		Uri:      uri,
	}
}

func (m *MongoDB) Connect(ctx context.Context) (client *mongo.Client, cleanup func()) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.getFormatUri()))
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	res, err := client.ListDatabaseNames(ctx, bson.D{})
	fmt.Println(res)

	return client, func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
		fmt.Println("mongodb disconnected")
	}
}

func (m *MongoDB) getFormatUri() string {
	return fmt.Sprintf(m.Uri, m.Username, m.Password)
}
