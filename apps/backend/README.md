# Backend – Go + PostgreSQL

RESTful API server built with Go (Gin) and PostgreSQL.

## Tech Stack

- **Go** 1.21+
- **Gin** – HTTP web framework
- **pgx v5** – PostgreSQL driver
- **godotenv** – Environment variable loading

## Project Structure

```
apps/backend/
├── cmd/
│   └── server/
│       └── main.go         # Application entry point
├── internal/
│   ├── config/             # Configuration loading
│   ├── database/           # DB connection & migrations
│   ├── handlers/           # HTTP route handlers
│   ├── middleware/         # CORS, logging, auth middleware
│   ├── models/             # Data models & response types
│   ├── repositories/       # Data access layer (SQL queries)
│   └── services/           # Business logic layer
├── migrations/             # SQL migration files
├── .env.example            # Environment variable template
├── Dockerfile              # Container image
└── go.mod                  # Go module definition
```

## Quick Start

### Prerequisites

- Go 1.21+
- PostgreSQL 16+ (or Docker)

### 1. Setup Environment

```bash
cp .env.example .env
# Edit .env with your database credentials
```

### 2. Start PostgreSQL (Docker)

```bash
# From the workspace root
docker-compose up postgres -d
```

### 3. Run the Server

```bash
# From workspace root using Nx
npm run dev:backend

# Or directly with Go
cd apps/backend
go run ./cmd/server
```

The server starts on `http://localhost:8080` by default.

## API Endpoints

| Method | Endpoint          | Description          |
|--------|-------------------|----------------------|
| GET    | `/api/health`     | Health check         |
| GET    | `/api/users`      | List all users       |
| GET    | `/api/users/:id`  | Get user by ID       |
| POST   | `/api/users`      | Create user          |
| PUT    | `/api/users/:id`  | Update user          |
| DELETE | `/api/users/:id`  | Delete user          |

### Health Check Response

```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T00:00:00Z",
  "version": "1.0.0",
  "services": {
    "api": "healthy",
    "database": "healthy"
  }
}
```

### Create User Request

```json
{
  "email": "user@example.com",
  "name": "John Doe"
}
```

## Building

```bash
# Build binary
go build -o dist/server ./cmd/server

# Build Docker image
docker build -t white-label-backend .
```

## Testing

```bash
go test ./...
```

## Environment Variables

| Variable              | Default           | Description                  |
|-----------------------|-------------------|------------------------------|
| `APP_ENV`             | `development`     | Application environment      |
| `BACKEND_PORT`        | `8080`            | HTTP server port             |
| `DB_HOST`             | `localhost`       | PostgreSQL host              |
| `DB_PORT`             | `5432`            | PostgreSQL port              |
| `DB_NAME`             | `whitelabel_db`   | Database name                |
| `DB_USER`             | `postgres`        | Database user                |
| `DB_PASSWORD`         | `postgres`        | Database password            |
| `DB_SSL_MODE`         | `disable`         | PostgreSQL SSL mode          |
| `JWT_SECRET`          | –                 | JWT signing secret           |
| `JWT_EXPIRATION`      | `24h`             | JWT token expiration         |
| `CORS_ALLOWED_ORIGINS`| `http://localhost:3000` | Allowed CORS origins   |
