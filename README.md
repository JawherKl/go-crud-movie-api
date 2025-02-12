# 🎬 Go CRUD Movies API Example

**A modern and robust CRUD (Create, Read, Update, Delete) API for managing movies, built with Go and GORM, featuring JWT authentication, rate limiting, and more.**

![Repository Size](https://img.shields.io/github/repo-size/JawherKl/go-crud-movie-api)
![Last Commit](https://img.shields.io/github/last-commit/JawherKl/go-crud-movie-api)
![Issues](https://img.shields.io/github/issues-raw/JawherKl/go-crud-movie-api)
![Forks](https://img.shields.io/github/forks/JawherKl/go-crud-movie-api)
![Stars](https://img.shields.io/github/stars/JawherKl/go-crud-movie-api)

![go-crud-movie-api](https://github.com/JawherKl/go-crud-movie-api/blob/main/go_crud.png)

## 📋 Table of Contents
- Features
- Getting Started
- Prerequisites
- Installation
- Running the Application
- API Endpoints
- Environment Variables
- Contributing
- License

## ✨ Features
- **CRUD Operations**: Manage movies with Create, Read, Update, and Delete operations.
- **JWT Authentication**: Secure API access with JSON Web Tokens.
- **Rate Limiting**: Control API usage and prevent abuse.
- **Pagination, Filtering, and Sorting**: Efficient movie listings.
- **Swagger/OpenAPI Documentation**: Explore the API effortlessly.
- **PostgreSQL**: Robust database support.
- **Environment Variable Configuration**: Easy database connection setup.
- **JSON-based API**: Seamless integration with other services.

## 🚀 Getting Started
Follow these instructions to set up and run the project on your local machine.

## 🛠 Prerequisites
- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Swagger](https://swagger.io/tools/swagger-ui/) for API documentation (optional)

## 📥 Installation
1. **Clone the repository:**
   ```sh
   git clone https://github.com/your-username/go-crud-movies-api.git
   cd go-crud-movies-api
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Set up the database:**
   Create a PostgreSQL database and configure the environment variables as described below.

4. **Set up environment variables:**
   Create a `.env` file in the project root with the following content:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_NAME=your_db_name
   DB_PASSWORD=your_db_password
   JWT_SECRET_KEY=your_secret_key
   ```

## 🏃 Running the Application

1. **Run the project:**
   ```sh
   go run main.go
   ```

2. The API will be running at:
   - `http://localhost:8080`

## 🔗 API Endpoints

### 🔑 Authentication

- **Login (Generate JWT Token):**
  ```http
  POST /auth/login
  ```
  Request body:
  ```json
  {
    "username": "user",
    "password": "password"
  }
  ```
  Response:
  ```json
  {
    "token": "your_jwt_token"
  }
  ```

### 🎥 Movie Operations

- **Create a movie:**
  ```http
  POST /movies
  ```
  Request body:
  ```json
  {
    "name": "Movie Name",
    "description": "Movie Description"
  }
  ```
  Response:
  ```json
  {
    "id": 1,
    "name": "Movie Name",
    "description": "Movie Description"
  }
  ```

- **Get all movies with pagination and sorting:**
  ```http
  GET /movies?page=1&limit=10&sort=desc
  ```
  Response:
  ```json
  [
    {
      "id": 1,
      "name": "Movie Name",
      "description": "Movie Description"
    }
  ]
  ```

- **Get a single movie by ID:**
  ```http
  GET /movies/{id}
  ```
  Response:
  ```json
  {
    "id": 1,
    "name": "Movie Name",
    "description": "Movie Description"
  }
  ```

- **Update a movie by ID:**
  ```http
  PUT /movies/{id}
  ```
  Request body:
  ```json
  {
    "name": "Updated Movie Name",
    "description": "Updated Movie Description"
  }
  ```
  Response:
  ```json
  {
    "id": 1,
    "name": "Updated Movie Name",
    "description": "Updated Movie Description"
  }
  ```

- **Delete a movie by ID:**
  ```http
  DELETE /movies/{id}
  ```
  Response:
  ```json
  {
    "message": "Movie deleted successfully"
  }
  ```

## ⚙️ Environment Variables

The application uses the following environment variables for configuration:

- **DB_HOST**: The database host (default: `localhost`)
- **DB_PORT**: The database port (default: `5432`)
- **DB_USER**: The database user
- **DB_NAME**: The database name
- **DB_PASSWORD**: The database password
- **JWT_SECRET_KEY**: The secret key for JWT token generation

## 🤝 Contributing

Contributions are welcome! Feel free to fork the repository, make changes, and submit pull requests.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⭐ Stargazers over time
[![Stargazers over time](https://starchart.cc/JawherKl/go-crud-movie-api.svg?variant=adaptive)](https://starchart.cc/JawherKl/go-crud-movie-api)
