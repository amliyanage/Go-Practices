# Go JWT Tasks API

A RESTful API for task management with JWT authentication built with Go, Gin, and GORM.

## Features

- 🔐 JWT-based authentication
- 👤 User registration and login
- ✅ CRUD operations for tasks
- 🔄 Live reloading with Air (development)
- 🐳 Docker and Docker Compose support
- 🗄️ MySQL database with GORM
- 🔥 Hot-reload development environment

## Tech Stack

- **Go 1.23+**
- **Gin** - HTTP web framework
- **GORM** - ORM library
- **MySQL** - Database
- **JWT** - Authentication
- **Air** - Live reloading
- **Docker** - Containerization

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Docker and Docker Compose
- Make

### Installation

1. Clone the repository:

```bash
git clone <your-repo-url>
cd go-jwt-tasks
```

2. Copy `.env.example` to `.env` and update the values:

```bash
cp .env.example .env
```

3. Start the application with Docker Compose:

```bash
make up
```

4. View logs:

```bash
make logs
```

5. Stop the application:

```bash
make down
```

## API Endpoints

### Health Check

```
GET /health
```

### Authentication

#### Register

```
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Login

```
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

Response:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "created_at": "2025-12-10T10:00:00Z",
    "updated_at": "2025-12-10T10:00:00Z"
  }
}
```

### User Profile (Protected)

#### Get Profile

```
GET /api/v1/profile
Authorization: Bearer <token>
```

### Tasks (Protected)

All task endpoints require authentication. Include the JWT token in the Authorization header:

```
Authorization: Bearer <your-token>
```

#### Create Task

```
POST /api/v1/tasks
Content-Type: application/json

{
  "title": "Complete project",
  "description": "Finish the Go JWT Tasks API"
}
```

#### Get All Tasks

```
GET /api/v1/tasks
```

Response:

```json
[
  {
    "id": 1,
    "title": "Complete project",
    "description": "Finish the Go JWT Tasks API",
    "completed": false,
    "user_id": 1,
    "created_at": "2025-12-10T10:00:00Z",
    "updated_at": "2025-12-10T10:00:00Z"
  }
]
```

#### Get Task by ID

```
GET /api/v1/tasks/:id
```

#### Update Task

```
PUT /api/v1/tasks/:id
Content-Type: application/json

{
  "title": "Updated title",
  "description": "Updated description",
  "completed": true
}
```

Note: All fields are optional. Only send the fields you want to update.

#### Delete Task

```
DELETE /api/v1/tasks/:id
```

## Development

### Available Make Commands

```bash
make help           # Show available commands
make build          # Build Docker image
make run            # Run Docker container
make dev            # Run with Air hot-reload (development mode)
make dev-stop       # Stop development container
make up             # Docker compose up
make down           # Docker compose down
make logs           # View docker compose logs (follow mode)
make clean          # Remove Docker container and image
```

### Running Locally (without Docker)

1. Install dependencies:

```bash
go mod download
```

2. Install Air for live reloading:

```bash
go install github.com/air-verse/air@v1.61.4
```

3. Start MySQL database (or use Docker):

```bash
docker run --name mysql -e MYSQL_ROOT_PASSWORD=yourpassword -e MYSQL_DATABASE=tasks_db -p 3306:3306 -d mysql:8.0
```

4. Update `.env` file with your database credentials

5. Run the application with Air:

```bash
air
```

Or run without Air:

```bash
go run main.go
```

## Environment Variables

Create a `.env` file in the root directory:

```env
DB_USER=root
DB_PASSWORD=sanmark@1234
DB_HOST=localhost
DB_PORT=3306
DB_NAME=tasks_db
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
PORT=8080
```

## Project Structure

```
.
├── config/          # Configuration management
├── controllers/     # HTTP handlers
│   ├── auth.go     # Authentication controllers
│   └── task.go     # Task controllers
├── middleware/      # Middleware functions
│   └── auth.go     # JWT authentication middleware
├── models/          # Database models
│   ├── user.go     # User model
│   └── task.go     # Task model
├── repo/            # Database connection
│   └── db.go       # Database initialization
├── docker/          # Docker files
│   ├── Dockerfile
│   ├── Dockerfile.dev
│   └── docker-compose.yml
├── .air.toml        # Air configuration
├── .env             # Environment variables
├── main.go          # Application entry point
├── Makefile         # Make commands
└── README.md        # This file
```

## Testing with cURL

### Register a new user

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"john_doe","email":"john@example.com","password":"password123"}'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'
```

### Create a task (replace TOKEN with your JWT)

```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer TOKEN" \
  -d '{"title":"My Task","description":"Task description"}'
```

### Get all tasks

```bash
curl -X GET http://localhost:8080/api/v1/tasks \
  -H "Authorization: Bearer TOKEN"
```

## License

MIT

## Author

Your Name
