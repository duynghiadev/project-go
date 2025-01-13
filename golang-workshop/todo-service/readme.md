# call api

All api in folder `todo-service` in Postman Applications

---

[post reference](https://hackernoon.com/how-to-build-your-own-todo-list-service-with-golang-and-mongodb)

---

# TODO List Microservice with MongoDB

Structure:

```
todo-list/
│
├── cmd/
│   └── main.go
├── pkg/
│   └── handler
│      └── add_task.go
│      └── http_handler.go
│      └── complite_task.go
│      └── get_all_task.go
│       └── update_task.go
│   └── mapper
│       └── task.go
│   └── model
│        └── task.go
│   └── usecase
│        └── task
│           └── repository
│               └── add_task.go
│               └── complite_task.go
│               └── get_all_task.go
│               └── mongo_repositiry.go
│               └── repository.go
│               └── update_task.go
│           └── service
│               └── add_task.go
│               └── complite_task.go
│               └── get_all_task.go
│               └── service.go
│               └── update_task.go
└── go.mod
```

## Overview

This presentation outlines the creation of a simple TODO list microservice using Golang and MongoDB. The service provides four main endpoints to manage tasks: adding a new task, updating a task by its ID, marking a task as completed, and retrieving all tasks.

## Components

1. **Main Application**: Initializes the MongoDB connection and sets up HTTP routes.
2. **Handler Layer**: Manages the HTTP request and response logic.
3. **Service Layer**: Contains business logic and interacts with the repository layer.
4. **Repository Layer**: Interacts directly with MongoDB to perform CRUD operations.

## Main Application

The main application is responsible for:

- Establishing a connection to the MongoDB database.
- Initializing repositories, services, and handlers.
- Setting up HTTP routes to handle incoming requests.

## Handler Layer

The handler layer:

- Receives HTTP requests.
- Validates and parses request data.
- Calls the appropriate service methods.
- Returns appropriate HTTP responses.

### Endpoints

1. **Add Task**: Adds a new task to the TODO list.

   - **Method**: POST
   - **URL**: `/add`
   - **Request Body**: JSON containing task details (e.g., title, completed status).

2. **Update Task**: Updates an existing task by its ID.

   - **Method**: PUT
   - **URL**: `/api/v1/update?id=<task_id>`
   - **Request Body**: JSON containing the fields to be updated (e.g., title).

3. **Complete Task**: Marks a task as completed by its ID.

   - **Method**: POST
   - **URL**: `/complete?id=<task_id>`

4. **Get All Tasks**: Retrieves all tasks from the TODO list.
   - **Method**: GET
   - **URL**: `/get-all`

## Service Layer

The service layer:

- Contains business logic for managing tasks.
- Interacts with the repository layer to perform database operations.

### Service Methods

1. **AddTask**: Adds a new task to the database.
2. **UpdateTask**: Updates the title of an existing task.
3. **CompleteTask**: Marks a task as completed.
4. **GetAllTasks**: Retrieves all tasks from the database.

## Repository Layer

The repository layer:

- Directly interacts with MongoDB to perform CRUD operations.
- Uses MongoDB's Go driver to connect and execute queries.

### Database Operations

1. **AddTask**: Inserts a new task document into the tasks collection.
2. **UpdateTask**: Updates the title of a task document based on its ObjectID.
3. **CompleteTask**: Updates the completed status of a task document.
4. **GetAllTasks**: Retrieves all task documents from the tasks collection.

## Key Points

- **MongoDB Integration**: The service connects to MongoDB using the official MongoDB Go driver.
- **UUID for Task IDs**: MongoDB's ObjectID is used as the unique identifier for tasks, allowing MongoDB to auto-generate IDs.
- **Layered Architecture**: The microservice is structured into distinct layers (handler, service, repository) to separate concerns and improve maintainability.
- **HTTP Endpoints**: The service provides HTTP endpoints to interact with the TODO list, supporting common CRUD operations.
