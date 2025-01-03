# Go CRUD Movies API Example

This is a simple and modern example of a CRUD (Create, Read, Update, Delete) API for managing movies, built using Go and GORM (Go Object Relational Mapper), along with added features such as JWT authentication, rate limiting, pagination, and Swagger documentation.

## Table of Contents
- [Features](#features)
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Features
- **Create, Read, Update, and Delete (CRUD) operations** for managing movies.
- **JWT-based Authentication** for secured API access.
- **Rate Limiting** to control API usage and prevent abuse.
- **Pagination, Filtering, and Sorting** for movie listings.
- **Swagger/OpenAPI Documentation** for easy API exploration.
- **Uses PostgreSQL** as the database.
- **Environment variable configuration** for database connection.
- Simple **JSON-based API** for seamless integration.

## Getting Started
Follow these instructions to set up and run the project on your local machine.

## Prerequisites
- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Swagger](https://swagger.io/tools/swagger-ui/) for documentation (optional)

## Installation
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
   - Create a PostgreSQL database and configure the environment variables as described below.

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

## Running the Application

1. **Run the project:**
   ```sh
   go run main.go
   ```

2. The API will be running at:
   - `http://localhost:8080`

## API Endpoints

### Authentication

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

### Movie Operations

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

## Environment Variables

The application uses the following environment variables for configuration:

- **DB_HOST**: The database host (default: `localhost`)
- **DB_PORT**: The database port (default: `5432`)
- **DB_USER**: The database user
- **DB_NAME**: The database name
- **DB_PASSWORD**: The database password
- **JWT_SECRET_KEY**: The secret key for JWT token generation

## Contributing

Contributions are welcome! Feel free to fork the repository, make changes, and submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
