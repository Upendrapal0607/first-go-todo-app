package model

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

// Todo represents the structure of a todo item.
type Todo struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Title     string             `bson:"title"`
    Completed bool               `bson:"completed"`
}

// TodoRepository handles CRUD operations for todos.
type TodoRepository struct {
    Collection *mongo.Collection
}

// Create inserts a new todo into the database.
func (r *TodoRepository) Create(ctx context.Context, todo Todo) (primitive.ObjectID, error) {
    todo.ID = primitive.NewObjectID()
    _, err := r.Collection.InsertOne(ctx, todo)
    return todo.ID, err
}

// GetAll retrieves all todo items.
func (r *TodoRepository) GetAll(ctx context.Context) ([]Todo, error) {
    var todos []Todo
    cursor, err := r.Collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var todo Todo
        if err := cursor.Decode(&todo); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

// get by Id
func (r *TodoRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*Todo, error) {
    var todo Todo
    err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&todo)
    return &todo, err
}

// Update a todo item
func (r *TodoRepository) Update(ctx context.Context, id primitive.ObjectID, todo Todo) error {
    _, err := r.Collection.UpdateOne(
        ctx,
        bson.M{"_id": id},                                      
        bson.M{"$set": bson.M{"title": todo.Title, "completed": todo.Completed}},
    )
    return err
}

// Delete a todo by ID
func (r *TodoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
    _, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}