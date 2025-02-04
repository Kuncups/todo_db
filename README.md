# ToDo API

> Aplikasi ToDo menggunakan Golang & PostgreSQL.

INI ADALAH TECHNICAL TES

## Endpoints

| Method | Endpoint   | Description       |
| ------ | ---------- | ----------------- |
| POST   | /tasks     | Create a new task |
| GET    | /tasks     | Get all tasks     |
| GET    | /tasks/:id | Get task by ID    |
| PUT    | /tasks/:id | Update task       |
| DELETE | /tasks/:id | Delete task       |

## Example Request & Response

## POST /tasks

Request :

```sh
{
  "title": "Task Title",
  "description": "Task Description",
  "status": "pending",
  "due_date": "2025-02-10"
}
```

Response :

```sh
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

## GET /tasks

Request :

```sh
GET /tasks?status=pending&page=1&limit=10&search=meeting
```

Response :

```sh
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

Request :

```sh
GET /tasks/:id
```

Response :

```sh
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

## PUT /tasks/:id

Request :

```sh
{
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "status": "completed",
  "due_date": "2025-02-15"
}
```

Response :

```sh
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

## DELETE /tasks/:id

Request :

```sh
DELETE /tasks/1
```

Response :

```sh
{
  "message": "Task deleted successfully"
}
```

## Release History

- 0.1.0
  - The first proper release
- 0.0.1
  - Work in progress

## Meta

Muhammad Yusuf Danan Risdianto – [@myusufdananr](https://instagram.com/myusufdananr) – myusufdananr@gmail.com

Distributed under the XYZ license. See `LICENSE` for more information.

[https://github.com/Kuncups/todo_db.git](https://github.com/Kuncups/)

## Contributing

1. Fork it (<https://github.com/Kuncups/todo_db.git/fork>)
2. Create your feature branch (`git checkout -b feature/todo_api`)
3. Commit your changes (`git commit -am 'first commit'`)
4. Push to the branch (`git push origin feature/todo_api`)
5. Create a new Pull Request
