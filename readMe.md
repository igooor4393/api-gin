# Golang Gin-Gorm-Postgres API

This project is a RESTful API built with Golang, using the Gin web framework and Gorm ORM with a PostgreSQL database. It includes a structured setup with modular code organization for controllers, routes, and database migrations. The API allows clients to interact with `Program` entities, providing CRUD operations.

## Features

- **Modular Controllers**: The `PostController` encapsulates the logic for handling requests related to the `Program` entity.
- **CRUD Operations**: Supports creating, reading, updating, and deleting `Program` records.
- **Pagination**: Includes pagination support for listing `Program` records.
- **Error Handling**: Provides detailed error responses for various failure scenarios, such as validation errors and database conflicts.
- **Structured Responses**: Standardizes API responses with status codes and messages.

## Endpoints

The following are some of the endpoints provided by the `PostController` for operations on `Program` entities:

- `POST /api/programs`: Allows the creation of a new `Program`.

```
         ----------------- just for exemple -----------------
- `GET /api/programs`: Retrieves a paginated list of `Program` records.
- `GET /api/programs/:id`: Fetches a single `Program` record by its ID.
- `PUT /api/programs/:id`: Updates an existing `Program` record.
- `DELETE /api/programs/:id`: Deletes a `Program` record.
       ----------------------------------------------------
```

## Models

The `Program` model defines the structure for `Program` entities, including fields such as `Id`, `Name`, `NameEn`, `IsPublic`, and `ProjectID`.

## Setup

The project is containerized and can be set up locally by running:

```sh
docker-compose up
```
### Step for check app

1. http://localhost:8080/api/healthchecker - the link will show if the container and the group are working for the api
2. https://api.postman.com/collections/14775136-f59b3c23-2e88-45e3-92f4-925bfd864bf4?access_key=PMAT-01J0ZZTECR1RYJAG9DCBQA3Q0J - postman collection for testing(import this link in postman)
