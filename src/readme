# Movie API Server
## Table of Contents
- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)

## Introduction
This is a Go-based RESTful API server for managing movies. It includes endpoints for creating, reading, updating, and deleting movies. The project aims to provide a simple yet powerful tool for movie enthusiasts and developers alike.

## Installation
To install the dependencies and run the server:

### Prerequisites
- Go 1.18 or higher
- Git

### Steps
Clone the repository:
```bash
git clone https://github.com/SanjishMaharjan/Movie-CRUD-Go.git
```
```bash
cd Movie-CRUD-Go
```
Install the dependencies:

```bash
go mod download
go mod tidy
```

### Configure the database:
Create an SQLite database file or modify the src/database/init.go file to configure your preferred database.
Run the application:
```bash
go run main.go
```

The server will start on http://127.0.0.1:3000.

### Usage
## API Endpoints
### Get all movies
- GET /movies/
### Get a movie by ID
- GET /movies/:id
### Add a new movie
- POST /movies/add

### Request Body:
 ```json
{
  "title": "Inception",
  "director": "Christopher Nolan",
  "release_year": 2010
}
```
### Update an existing movie by ID
- PUT /movies/add/:id
### Request Body:
```json
{
  "title": "The Dark Knight",
  "director": "Christopher Nolan",
  "release_year": 2008
}
```
### Delete a movie by ID
```bash
DELETE /movies/delete/:id
```
## Project Structure
/src
├── /database
│   └── init.go         
├── /models
│   └── movie.go        
├── /handlers
│   ├── movie.go      
└── main.go            

