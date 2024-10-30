package mongodb

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
    Client *mongo.Client
}

// Initialize a new MongoDB client
func NewMongoClient(uri string) (*MongoClient, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        return nil, fmt.Errorf("MongoDB ping failed: %w", err)
    }

    fmt.Println("Connected to MongoDB")
    return &MongoClient{Client: client}, nil
}

// Close the MongoDB client
func (mc *MongoClient) Close() {
    if err := mc.Client.Disconnect(context.Background()); err != nil {
        fmt.Println("Failed to disconnect from MongoDB:", err)
    } else {
        fmt.Println("Disconnected from MongoDB")
    }
}
