package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
)

var client *mongo.Client

func Connect() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var err error
    client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        return err
    }

    return client.Ping(ctx, nil)
}

func GetCollection(name string) *mongo.Collection {
    return client.Database("Test").Collection(name)
}