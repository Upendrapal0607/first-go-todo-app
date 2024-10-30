package delete

import (
    "context"
    "go-basic-todo-api/module/todo/model"
	 "go.mongodb.org/mongo-driver/bson/primitive"
    "net/http"
    "github.com/go-chi/chi/v5"
)

type TodoHandler struct {
    Repo *model.TodoRepository
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
    idParam := chi.URLParam(r, "id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    if err := h.Repo.Delete(context.Background(), id); err != nil {
        http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
