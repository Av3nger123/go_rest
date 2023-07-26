package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	TodoCollection *mongo.Collection
	mongoConn      *mongo.Client
)

func Connect() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	mongoConn, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	TodoCollection = mongoConn.Database("todo-app").Collection("todos")
	return nil
}

func Disconnect() {
	if mongoConn != nil {
		mongoConn.Disconnect(context.Background())
	}
}
