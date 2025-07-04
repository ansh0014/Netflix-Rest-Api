# Netflix Watchlist API

A simple RESTful API built with Go and MongoDB for managing a Netflix watchlist.

## Project Overview

This project provides a backend API for tracking Netflix movies/shows that you want to watch or have already watched. It uses MongoDB Atlas as the database and implements a RESTful API with the following features:

- Get all movies in the watchlist
- Add a new movie to the watchlist
- Mark a movie as watched
- Delete a specific movie
- Delete all movies from the watchlist

## Project Structure

```
├── controller/
│   └── controller.go    # MongoDB connection and CRUD operations
├── model/
│   └── model.go         # Data model for Netflix watchlist items
├── router/
│   └── router.go        # API route definitions
├── main.go              # Application entry point
```

## Data Model

Each movie in the watchlist has the following structure:

```go
type Netflix struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	Watched bool   `json:"watched,omitempty" bson:"watched,omitempty"`
}
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/movies` | Get all movies in the watchlist |
| POST | `/api/movie` | Add a new movie to the watchlist |
| PUT | `/api/movie/{id}` | Mark a movie as watched |
| DELETE | `/api/movie/{id}` | Delete a specific movie |
| DELETE | `/api/deleteallmovie` | Delete all movies from the watchlist |

## How to Run

1. Make sure you have Go installed on your system
2. Clone this repository
3. Update the MongoDB connection string in `controller/controller.go` if needed
4. Run the application:
   ```
   go run main.go
   ```
5. The server will start at `http://localhost:9500`

## API Usage Examples

### Get all movies
```
GET http://localhost:9500/api/movies
```

### Add a new movie
```
POST http://localhost:9500/api/movie
Content-Type: application/json

{
  "name": "Stranger Things",
  "watched": false
}
```

### Mark a movie as watched
```
PUT http://localhost:9500/api/movie/{id}
```

### Delete a movie
```
DELETE http://localhost:9500/api/movie/{id}
```

### Delete all movies
```
DELETE http://localhost:9500/api/deleteallmovie
```