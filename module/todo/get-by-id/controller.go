package get_all

import (
    "context"
    "encoding/json"
    "net/http"
    "go-basic-todo-api/module/todo/model"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
    Repo *model.TodoRepository
}


func (h *TodoHandler) GetSingleTodo(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }
    todos, err := h.Repo.GetByID(context.Background(),id)
    if err != nil {
        http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(todos)
}
