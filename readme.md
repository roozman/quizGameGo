# Quiz Game Backend

A simple backend service for a multiple-choice quiz game, built in Go. This project is actively developed as a learning exercise to get comfortable with Go, SQL, and building RESTful APIs.

---

## Overview

The service provides user authentication (registration and login) and will eventually handle game logic, including categories, difficulty levels, and scoring. Currently, it manages user accounts with phone numbers and passwords.

---

## Features

- **User Registration** – create an account with name, phone number, and password.
- **User Login** – authenticate using phone number and password (MD5 hashed, to be improved).
- **Health Check** – simple endpoint to verify service availability.
- **Database** – MySQL containerized via Docker Compose.
- **Modular architecture** – separates HTTP handlers, business logic (service layer), and data persistence (repository pattern).

---

## Tech Stack

- **Go 1.21+** – standard `net/http` for routing.
- **MySQL 8.0** – relational database.
- **Docker & Docker Compose** – for local database setup.
- **`go-sql-driver/mysql`** – MySQL driver.

---

## Getting Started

### Prerequisites

- Go 1.21 or later
- Docker & Docker Compose (optional, for running MySQL)

### Clone the Repository

```bash
git clone https://github.com/yourusername/quiz-game-backend.git
cd quiz-game-backend
```

### Start the Database
A `docker-compose.yml` is provided to launch a MySQL container with the required schema.

```bash
docker-compose up -d
```
This will:
- Start MySQL on port `3308`
- Create the database `quizgame_db`
- Run the `/repository/mysql/setup_db.sql` to create the `users` table

### Build and Run the Server
```bash
go mod tidy
go run main.go
```
The server will listen on `http://localhost:8080`

---
## API Documentation

All endpoints return JSON.

### Health Check
```http request
GET /health-check
```
Response: `{"alive": "true"}`

### User Registration
```http request
POST /users/register
```
Request body:
```json
{
  "name": "your_name",
  "phone_number": "09000000000",
  "password": "your_password"
}
```
On success (200) `{"message":"user registered"}`

### User Login
```http request
POST /users/login
```
Request body:
```json
{
  "phone_number": "09000000000",
  "password": "your_password"
}
```
On success (200) `{"message":"user credentials are ok"}`

Note: All error responses include a descriptive errors field.