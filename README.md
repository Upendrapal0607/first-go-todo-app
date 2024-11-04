# Todo API

This is a simple Todo API built with Go. The API allows users to create, read, update, and delete todos.

## Clone Repository

To clone the repository, run the following command:

```bash
git clone <https://github.com/Upendrapal0607/first-go-todo-app.git>
```

## Setup

1. Change to the project directory:
   ```bash
   cd first-go-todo-app
   ```
 
2. Install the necessary Go packages (if you have a `go.mod` file):
   ```bash
   go mod tidy
   ```

## Start Server

To start the server, run the following command:

```bash
go run server.go
```

The server will run at `http://localhost:8080`.

## API Endpoints

### Available Routes
- **POST /todos**  
  Create a new todo item.  
 
 ```
 /api/todo/v1/todos
 ```

- **GET /todos**  
  Retrieve all todo items.  
  ```
  /api/todo/v1/todos
  ```

- **GET /todos/{id}**  
  Retrieve a single todo item by ID.  
  ```
  /api/todo/v1/todos/{id}
  ```


- **PUT /todos/{id}**  
  Update an existing todo item by ID.  

  ```
  /api/todo/v1/todos/{id}
  ```
- **DELETE /todos/{id}**  
  Delete a todo item by ID.  
  ```
  /api/todo/v1/todos/{id}
  ```

## Example

To create a new todo item, you can send a POST request to `http://localhost:8080/api/todo/v1/todos` with the required data in the request body.

## License

