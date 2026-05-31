# Auth Service

Authentication microservice built with Go, Gin, PostgreSQL, JWT, and Docker.

## Features

* User Registration
* User Authentication
* JWT Token Generation
* Protected Endpoints
* PostgreSQL Persistence
* Docker Support
* Clean Architecture
* Repository Pattern

## Technology Stack

* Go 1.24
* Gin Framework
* PostgreSQL
* JWT
* bcrypt
* Docker
* Docker Compose

## Running Locally

### Start PostgreSQL

```bash
docker compose up -d postgres
```

### Run Service

```bash
go run ./cmd/server
```

Service:

```text
http://localhost:8080
```

## API Endpoints

### POST /auth/register

Creates a new user account.

### POST /auth/login

Authenticates a user and returns a JWT access token.

### GET /me

Returns authenticated user information.

Requires:

```text
Authorization: Bearer <token>
```

## Architecture

```text
Request
  ↓
Handler
  ↓
Service
  ↓
Repository
  ↓
PostgreSQL
```

## Future Improvements

* Refresh Tokens
* Role Based Access Control (RBAC)
* OAuth2
* OpenTelemetry
* Prometheus Metrics
* Kubernetes Deployment
