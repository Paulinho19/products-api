# Products API

Simple REST API for managing products, built with Go and [Gin](https://github.com/gin-gonic/gin).

## Tech Stack

- Go 1.26
- Gin (HTTP framework)
- PostgreSQL (via Docker)

## Prerequisites

- Go 1.26+
- Docker / Docker Compose

## Getting Started

1. Start the database:

   ```bash
   docker-compose up -d
   ```

2. Run the API:

   ```bash
   go run cmd/main.go
   ```

The server starts on `http://localhost:8000`.

## Endpoints

| Method | Route       | Description          |
|--------|-------------|----------------------|
| GET    | `/ping`     | Health check         |
| GET    | `/products` | List all products    |

### Example

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

## Project Structure

```
cmd/            entrypoint (main.go)
controller/     HTTP handlers
model/          domain models
docker-compose.yml  PostgreSQL container
```

## Database

PostgreSQL runs via Docker Compose:

- Host: `localhost:5432`
- User: `postgres`
- Password: `1234`
- Database: `products`
