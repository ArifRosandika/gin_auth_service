Gin Auth Service

A backend authentication service built with Go (Gin) implementing JWT-based authentication with access token & refresh token lifecycle, Redis-backed refresh token storage, and PostgreSQL for persistent user data.
The project follows Clean Architecture principles to keep business logic isolated, testable, and maintainable.

![GitHub last commit](https://img.shields.io/github/last-commit/ArifRosandika/gin_auth_service?color=blue)
![GitHub repo size](https://img.shields.io/github/repo-size/ArifRosandika/gin_auth_service)
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-00ADD8?logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-DC382D?logo=redis&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?logo=jsonwebtokens&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)

📘 Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Authentication Flow](#authentication-flow)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Run with Docker](#run-with-docker)
  - [Run Locally](#run-locally)
- [Environment Variables](#environment-variables)
- [API Testing](#api-testing)
- [License](#license)

✨ Features

User registration & login

JWT Access Token authentication

Refresh Token lifecycle management

Refresh Token revocation (logout & rotation)

Redis as authoritative refresh token store

PostgreSQL persistence with GORM

Clean Architecture (Handler → Usecase → Repository)

Environment-based configuration using Viper

Dockerized with PostgreSQL & Redis via Docker Compose

🧱 Tech Stack
Backend

Go

Gin – HTTP framework

GORM – ORM for PostgreSQL

JWT (HS256) – Authentication

Redis – Refresh token storage & revocation

Argon2id – Password hashing

Infrastructure

PostgreSQL

Redis

Docker & Docker Compose

🔐 Authentication Flow
Login

Validate user credentials

Generate Access Token (short-lived)

Generate Refresh Token

Store refresh token in Redis

refresh:<token> -> user_id

Refresh Token

Client sends refresh token

Server validates token existence in Redis

Issue new access token

Revoke old refresh token (delete from Redis)

Logout

Client sends refresh token

Refresh token is deleted from Redis

Token becomes unusable immediately

Refresh tokens are stateful and fully controlled by Redis.

📁 Project Structure
gin_auth_service/
├── cmd/
│   └── main.go                # Application entry point
│
├── config/
│   └── database.go            # Database initialization
│
├── internal/
│   ├── cache/
│   │   └── redis.go            # Redis cache wrapper
│   │
│   ├── delivery/
│   │   └── http/
│   │       ├── dto/
│   │       │   ├── request/    # HTTP request DTOs
│   │       │   └── response/   # HTTP response DTOs
│   │       ├── handler/        # HTTP handlers (controllers)
│   │       ├── middleware/     # HTTP middlewares (JWT, etc)
│   │       └── router/         # Route definitions
│   │
│   ├── domain/
│   │   ├── auth_usecase_interface.go
│   │   └── user_usecase_interface.go
│   │
│   ├── repository/
│   │   ├── interfaces/
│   │   │   ├── redis_token_repository_interface.go
│   │   │   └── user_repository_interface.go
│   │   └── impl/
│   │       ├── redis_token_repository.go
│   │       └── user_repository.go
│   │
│   └── usecase/
│       ├── auth_usecase.go
│       ├── token_usecase.go
│       └── user_usecase.go
│
├── env/
│   └── .env.example            # Environment variables template
│
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md


🚀 Getting Started
Prerequisites

Docker & Docker Compose

Go 1.22+

Run with Docker
docker-compose up --build

Run Locally (without Docker)
go mod tidy
go run cmd/main.go

🔧 Environment Variables
# App
APP_PORT=8080
JWT_SECRET=your_secret_key

# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=auth_service_db

# Redis
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=

🧪 API Testing

A test.rest file is included for:

Register

Login

Profile

Refresh token

Logout

Compatible with VS Code REST Client extension.

📜 License

This project is licensed under the MIT License.
