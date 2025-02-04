# ToDo API

Aplikasi ToDo menggunakan Golang & PostgreSQL.

## INI ADALAH TECHNICAL TES

1. **Clone repository**

   ```sh
   git clone https://github.com/Kuncups/todo-api.git
   cd todo-api

   ```

## Endpoints

| Method | Endpoint   | Description       |
| ------ | ---------- | ----------------- |
| POST   | /tasks     | Create a new task |
| GET    | /tasks     | Get all tasks     |
| GET    | /tasks/:id | Get task by ID    |
| PUT    | /tasks/:id | Update task       |
| DELETE | /tasks/:id | Delete task       |

## Example Request & Response

### POST /tasks

**Request Body**:

```json
{
  "title": "Task Title",
  "description": "Task Description",
  "status": "pending",
  "due_date": "2025-02-10"
}
```

**Response**

```json
{
  "message": "Task created successfully",
  "task": {
    "id": 1,
    "title": "Task Title",
    "description": "Task Description",
    "status": "pending",
    "due_date": "2025-02-10"
  }
}
```

### GET /tasks

## Query Parameters:

1. status (optional): Filter tasks by status (e.g. completed, pending).
2. page (optional): Page number for pagination.
3. limit (optional): Number of tasks per page.
4. search (optional): Search term to filter tasks by title or description.

**Example Request**:

```json
GET /tasks?status=pending&page=1&limit=10&search=meeting
```

**Response**

```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Task Title",
      "description": "Task Description",
      "status": "pending",
      "due_date": "2025-02-10"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_tasks": 1
  }
}
```

```json
GET /tasks/:id
```

**Response**

```json
{
  "message": "Task retrieved successfully",
  "task": {
    "id": 1,
    "title": "Task Title",
    "description": "Task Description",
    "status": "pending",
    "due_date": "2025-02-10"
  }
}
```

### PUT /tasks/:id

**Request Body**:

```json
{
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "status": "completed",
  "due_date": "2025-02-15"
}
```

**Response**

```json
{
  "message": "Task updated successfully",
  "task": {
    "id": 1,
    "title": "Updated Task Title",
    "description": "Updated Task Description",
    "status": "completed",
    "due_date": "2025-02-15"
  }
}
```

### DELETE /tasks/:id

**Example Request**:

```json
DELETE /tasks/1
```

**Response**

```json
{
  "message": "Task deleted successfully"
}
```
#   t o d o _ d b  
 