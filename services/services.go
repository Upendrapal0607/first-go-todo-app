package services

import (
    "fmt"
    "go-basic-todo-api/utils/mongodb"
)

func InitMongoClient() (*mongodb.MongoClient, error) {
    uri := os.Getenv("DEFAULT_MONGODB_URL")
    if uri == "" {
        return nil, fmt.Errorf("DEFAULT_MONGODB_URL not set in environment")
    }

    mongoClient, err := mongodb.NewMongoClient(uri)
    if err != nil {
        return nil, err
    }

    return mongoClient, nil
}
