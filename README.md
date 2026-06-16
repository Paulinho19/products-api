# Products API

REST API for managing products, built with Go, Gin, and PostgreSQL.

## Tech Stack

| Tool | Version |
|------|---------|
| Go | 1.26.2 |
| Gin | v1.12.0 |
| lib/pq | v1.12.3 |
| PostgreSQL | 12 (Docker) |

## Prerequisites

- Go 1.26+
- Docker / Docker Compose

## Getting Started

1. Start the database:

   ```bash
   docker-compose up -d
   ```

2. Create the table:

   ```sql
   CREATE TABLE product (
     id           SERIAL PRIMARY KEY,
     product_name VARCHAR(100) NOT NULL,
     price        NUMERIC(10, 2) NOT NULL
   );
   ```

3. Run the API:

   ```bash
   go run cmd/main.go
   ```

Server starts on `http://localhost:8000`.

## Endpoints

| Method | Route       | Description       |
|--------|-------------|-------------------|
| GET    | `/ping`     | Health check      |
| GET    | `/products` | List all products |

### GET /ping

```bash
curl http://localhost:8000/ping
```

```json
{ "message": "pong" }
```

### GET /products

```bash
curl http://localhost:8000/products
```

```json
[
  {
    "product_id": 1,
    "name": "Máquina de Lavar Louças",
    "price": 1750.00
  }
]
```

## Architecture

Request flows through three layers:

```
HTTP Request
    ↓
Controller   (controller/)   — parse request, write response
    ↓
UseCase      (usecase/)      — business logic
    ↓
Repository   (repository/)   — SQL queries
    ↓
PostgreSQL
```

## Project Structure

```
cmd/
  main.go              entrypoint, router setup
controller/
  product_controller.go  HTTP handlers
usecase/
  product_usecase.go   business logic
repository/
  product_repository.go  database queries
model/
  product.go           domain structs
db/
  conn.go              database connection
docker-compose.yml     PostgreSQL container
```

## Database

PostgreSQL via Docker Compose:

| Setting  | Value      |
|----------|------------|
| Host     | localhost  |
| Port     | 5432       |
| User     | postgres   |
| Password | 1234       |
| Database | products   |
