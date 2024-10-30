package update_controller

import (
    "context"
    "encoding/json"
    "net/http"
    "go.mongodb.org/mongo-driver/bson/primitive"
	 "go-basic-todo-api/module/todo/model"
	 "github.com/go-chi/chi/v5"

)

type TodoHandler struct {
    Repo *model.TodoRepository
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
    idParam := chi.URLParam(r, "id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    var todo model.Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := h.Repo.Update(context.Background(), id, todo); err != nil {
        http.Error(w, "Failed to update todo", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}