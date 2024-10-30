// package create_controller

// import (
//     "context"
//     "encoding/json"
//     "net/http"
// 	"go-basic-todo-api/module/todo"
// )

// type TodoHandler struct {
//     Repo *todo.TodoRepository
// }

// func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
//     var todo todo.Todo
//     if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
//         http.Error(w, "Invalid input", http.StatusBadRequest)
//         return
//     }

//     id, err := h.Repo.Create(context.Background(), todo)
//     if err != nil {
//         http.Error(w, "Failed to create todo", http.StatusInternalServerError)
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
// }

package create

import (
    "context"
    "encoding/json"
    "net/http"
    "go-basic-todo-api/module/todo/model"
)

// TodoHandler handles todo operations.
type TodoHandler struct {
    Repo *model.TodoRepository
}

// CreateTodo creates a new todo item.
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
    var newTodo model.Todo
    if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    id, err := h.Repo.Create(context.Background(), newTodo)
    if err != nil {
        http.Error(w, "Failed to create todo", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}
