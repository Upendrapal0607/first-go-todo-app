package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "github.com/go-chi/chi/v5"
    s "go-basic-todo-api/services"
     "go-basic-todo-api/module/todo"
    "go-basic-todo-api/module/todo/model"

)

func main() {
    fmt.Println("Starting")


    // Initialize MongoDB client
    mongoClient, err := s.InitMongoClient()
    if err != nil {
        log.Fatalf("[mongodb] Error initializing MongoDB: %s", err)
    }
    defer mongoClient.Close()

    // Get the todos collection from the database
    collection := mongoClient.Client.Database("todoDB").Collection("todos")

    // Initialize repository and handler

    todoRepo := &model.TodoRepository{Collection: collection}
 

    fmt.Println("Connected to MongoDB: Database(todoDB), Collection(todos)")

    // Setup router
    r := chi.NewRouter()
    r.Get("/health/status", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "ok")
    })


    // todo route 
    r.Mount("/api/todo/v1",todo.Routes(todoRepo))
  
    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

    go func() {
        fmt.Println("Server running on http://localhost:8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()

    <-quit
    fmt.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    fmt.Println("Server stopped gracefully.")
}
