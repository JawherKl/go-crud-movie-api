# Go CRUD Movies API Example
This is a simple example of a CRUD (Create, Read, Update, Delete) API for managing movies, built using Go and GORM (Go Object Relational Mapper).

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
- Create, Read, Update, and Delete movies.
- Uses PostgreSQL as the database.
- Environment variable configuration for database connection.
- Simple JSON-based API.

## Getting Started
Follow these instructions to set up and run the project on your local machine.

## Prerequisites
- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation
1. **Clone the repository:**

   ```sh
   git clone https://github.com/your-username/go-crud-movies-api-example.git
   cd go-crud-movies-api-example
## Install dependencies:
go mod tidy

## Set up the database:
## Create a PostgreSQL database and configure the environment variables as described below.
Running the Application

## Set up environment variables:
## Create a .env file in the project root with the following content:
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_NAME=your_db_name
DB_PASSWORD=your_db_password
Run the application:

## Run the project
go run main.go
The API will be running at:
http://localhost:8080


## API Endpoints
Create a movie:
POST /movies
{
  "name": "Movie Name",
  "description": "Movie Description"
}

Get all movies:
GET /movies

Get a single movie by ID:
GET /movies/{id}

Update a movie by ID:
PUT /movies/{id}
{
  "name": "Updated Movie Name",
  "description": "Updated Movie Description"
}

Delete a movie by ID:
DELETE /movies/{id}

## Environment Variables
The application uses the following environment variables for configuration:

DB_HOST: The database host (default: localhost)
DB_PORT: The database port (default: 5432)
DB_USER: The database user
DB_NAME: The database name
DB_PASSWORD: The database password
