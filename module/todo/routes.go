package todo

import (
    "github.com/go-chi/chi/v5"
    "go-basic-todo-api/module/todo/model"
     allTodo "go-basic-todo-api/module/todo/get-all"
	 singleTodo "go-basic-todo-api/module/todo/get-by-id"
    create "go-basic-todo-api/module/todo/create"
    update "go-basic-todo-api/module/todo/update"
    delete "go-basic-todo-api/module/todo/delete"
)

// type TodoHandler struct {
//     Repo *TodoRepository
// }

// Routes sets up the todo routes.
func Routes(todoRepo *model.TodoRepository) chi.Router {
    r := chi.NewRouter()

    getTodoHandler := &allTodo.TodoHandler{Repo: todoRepo}
    getTodoByIdHandler := &singleTodo.TodoHandler{Repo: todoRepo}
    createTodoHandler := &create.TodoHandler{Repo: todoRepo}
    updateTodoHandler := &update.TodoHandler{Repo: todoRepo}
    deleteTodoHandler := &delete.TodoHandler{Repo: todoRepo}

    r.Post("/todos", createTodoHandler.CreateTodo)
    r.Get("/todos", getTodoHandler.GetAllTodos)
    r.Get("todos/{id}", getTodoByIdHandler.GetSingleTodo)
    r.Put("/todos/{id}", updateTodoHandler.UpdateTodo)
    r.Delete("/todos/{id}", deleteTodoHandler.DeleteTodo)

    return r
}
