# Auth-as-a-Service (AaaS)

A modern fullstack authentication microservice system inspired by Auth0 — built with Golang, MongoDB, RabbitMQ, and React.js.

## Overview

This project implements a secure, scalable user authentication system with JWT tokens, session management, and a React frontend. It demonstrates microservice architecture patterns using Go for backend services.

## Features

- User registration and authentication
- JWT and refresh token implementation
- Password encryption with bcrypt
- MongoDB for data persistence
- RabbitMQ for service communication
- Protected API routes
- React.js frontend

## Tech Stack

- **Backend**: Golang with standard HTTP package/mux
- **Frontend**: React.js with Tailwind CSS
- **Database**: MongoDB
- **Messaging**: RabbitMQ
- **Authentication**: JWT
- **Containerization**: Docker and Docker Compose

## Project Structure

```
auth-as-a-service/
├── backend/
│   ├── user-service/       # User management and authentication
│   ├── token-service/      # Token generation and validation
│   ├── session-service/    # Session tracking
│   └── common/             # Shared code and utilities
├── frontend/               # React.js application
├── docker/                 # Docker configuration
├── scripts/                # Helper scripts
└── docs/                   # Documentation
```

## Getting Started

### Prerequisites

- Go 1.18 or later
- Node.js 16 or later
- MongoDB
- RabbitMQ
- Docker and Docker Compose (optional)

### Backend Setup

```bash
# Clone the repository
git clone https://github.com/BhuvanMM/auth-as-a-service.git
cd auth-as-a-service

# Start the user service
cd backend/user-service
go mod tidy
go run main.go

# Start the token service (in a new terminal)
cd backend/token-service
go mod tidy
go run main.go

# Start the session service (in a new terminal)
cd backend/session-service
go mod tidy
go run main.go
```

### Frontend Setup

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start the React application
npm start
```

### Using Docker

```bash
# Start all services
docker-compose up

# Or build and start
docker-compose up --build
```

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Authenticate a user
- `POST /api/auth/refresh` - Refresh access token
- `POST /api/auth/logout` - Logout user

### User Management

- `GET /api/users/me` - Get current user profile
- `PUT /api/users/me` - Update user profile
- `PUT /api/users/password` - Change password

## Environment Variables

Create a `.env` file in each service directory with the following variables:

```
# MongoDB
MONGO_URI=mongodb://localhost:27017/auth_service

# JWT
JWT_SECRET=your_jwt_secret_key
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=7d

# RabbitMQ
RABBITMQ_URI=amqp://guest:guest@localhost:5672/

# Server
PORT=8080
ENV=development
```

## Testing

```bash
# Run backend tests
cd backend/user-service
go test ./...

# Run frontend tests
cd frontend
npm test
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
