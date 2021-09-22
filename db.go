package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getTable() (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("db"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			return nil, err
		}
	}()
	return client.Database("mainDB").Collection("paper"), nil
}
