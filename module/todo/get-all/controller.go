package get_all

import (
    "context"
    "encoding/json"
    "net/http"
    "go-basic-todo-api/module/todo/model"
)

type TodoHandler struct {
    Repo *model.TodoRepository
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
    todos, err := h.Repo.GetAll(context.Background())
    if err != nil {
        http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(todos)
}
