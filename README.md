
# Go TaskManager HTTP rest api

TaskManager is a simple task management application built with Golang, Gorilla Mux, and PostgreSQL. It allows you to create, edit, find, and delete tasks.

## Features

- Create a new task
- Edit an existing task
- Retrieve a list of all tasks
- Retrieve a single task by ID
- Delete a task

## Project Structure

```
taskmanager/
├── .env
├── go.mod
├── go.sum
├── main.go
├── models/
│   └── task.go
├── handlers/
│   └── task_handler.go
├── db/
│   └── db.go
├── validation/
│   └── validation.go
```

## Prerequisites

- Go 1.22.1 or higher
- PostgreSQL

## Packages Used

- [github.com/go-playground/validator/v10](https://github.com/go-playground/validator) - For data validation
- [github.com/gorilla/mux](https://github.com/gorilla/mux) - For HTTP routing
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - For loading environment variables from a `.env` file
- [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) - GORM PostgreSQL driver
- [gorm.io/gorm](https://github.com/go-gorm/gorm) - For ORM functionality

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/taskmanager.git
    cd taskmanager
    ```

2. Create a `.env` file in the root directory of your project and add the following environment variables:

    ```env
    DB_HOST=your_db_host
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    DB_PORT=your_db_port
    ```

3. Install the dependencies:

    ```sh
    go mod tidy
    ```

4. Run the application:

    ```sh
    go run main.go
    ```

## API Endpoints

### Create a Task

- **URL**: `/tasks`
- **Method**: `POST`
- **Request Body**:

    ```json
    {
        "title": "Task Title",
        "description": "Task Description",
        "completed": false
    }
    ```

- **Response**:

    ```json
    {
        "id": 1,
        "title": "Task Title",
        "description": "Task Description",
        "completed": false
    }
    ```

### Retrieve All Tasks

- **URL**: `/tasks`
- **Method**: `GET`
- **Response**:

    ```json
    [
        {
            "id": 1,
            "title": "Task Title",
            "description": "Task Description",
            "completed": false
        },
        ...
    ]
    ```

### Retrieve a Single Task

- **URL**: `/tasks/{id}`
- **Method**: `GET`
- **Response**:

    ```json
    {
        "id": 1,
        "title": "Task Title",
        "description": "Task Description",
        "completed": false
    }
    ```

### Update a Task

- **URL**: `/tasks/{id}`
- **Method**: `PUT`
- **Request Body**:

    ```json
    {
        "title": "Updated Task Title",
        "description": "Updated Task Description",
        "completed": true
    }
    ```

- **Response**:

    ```json
    {
        "id": 1,
        "title": "Updated Task Title",
        "description": "Updated Task Description",
        "completed": true
    }
    ```

### Delete a Task

- **URL**: `/tasks/{id}`
- **Method**: `DELETE`
- **Response**: `204 No Content`

## Contributing

Feel free to open issues or submit pull requests for new features, bug fixes, or improvements.