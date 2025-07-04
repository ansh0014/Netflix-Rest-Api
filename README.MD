# Netflix Watchlist API

[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![MongoDB](https://img.shields.io/badge/MongoDB-4.4+-47A248?style=flat&logo=mongodb)](https://www.mongodb.com/)

A RESTful API built with Go and MongoDB for managing a Netflix watchlist.

## 📋 Overview

This project provides a backend API for tracking Netflix movies/shows with the following features:

- Get all movies in the watchlist
- Add a new movie to the watchlist
- Mark a movie as watched
- Delete a specific movie
- Delete all movies from the watchlist

## 🛠️ Tech Stack

- **Go** - Programming language
- **MongoDB Atlas** - Database
- **Gorilla Mux** - HTTP router and URL matcher

## 📁 Project Structure

```
├── controller/  # MongoDB connection and CRUD operations
├── model/       # Data model for watchlist items
├── router/      # API route definitions
└── main.go      # Application entry point
```

## 🔌 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/movies` | Get all movies |
| POST | `/api/movie` | Add a new movie |
| PUT | `/api/movie/{id}` | Mark as watched |
| DELETE | `/api/movie/{id}` | Delete a movie |
| DELETE | `/api/deleteallmovie` | Delete all movies |

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- MongoDB Atlas account

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/yourusername/netflix-watchlist-api.git
   cd netflix-watchlist-api
   ```

2. Update the MongoDB connection string in `controller/controller.go`

3. Run the application
   ```bash
   go run main.go
   ```

4. The server will start at `http://localhost:9500`

## 📝 Usage Example

### Add a new movie
```bash
curl -X POST http://localhost:9500/api/movie \
  -H "Content-Type: application/json" \
  -d '{"name":"Stranger Things","watched":false}'
```
