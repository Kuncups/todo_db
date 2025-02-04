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
```

## GET /tasks

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
```

## Release History

- 0.1.0
  - The first proper release
  - CHANGE: Rename `foo()` to `bar()`
- 0.0.1
  - Work in progress

## Meta

Your Name – [@YourTwitter](https://twitter.com/dbader_org) – YourEmail@example.com

Distributed under the XYZ license. See `LICENSE` for more information.

[https://github.com/yourname/github-link](https://github.com/dbader/)

## Contributing

1. Fork it (<https://github.com/yourname/yourproject/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

<!-- Markdown link & img dfn's -->

[npm-image]: https://img.shields.io/npm/v/datadog-metrics.svg?style=flat-square
[npm-url]: https://npmjs.org/package/datadog-metrics
[npm-downloads]: https://img.shields.io/npm/dm/datadog-metrics.svg?style=flat-square
[travis-image]: https://img.shields.io/travis/dbader/node-datadog-metrics/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/dbader/node-datadog-metrics
[wiki]: https://github.com/yourname/yourproject/wiki
