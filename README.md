# Personal Book Library Manager
A Go-based web application for managing personal book collections with React frontend.

## Features
- Book management(CRUD operations)
- Book Status tracking(read/unread)
- Book rating (max 5)
- OpenLibrary API integration
- Search functionality
- RESTful API
- SQLite database

## Prerequisites
- Go 1.23.3 or higher
- SQLite3
- Node.js and npm (for frontend)

## Installation
1. Clone the repository

    ```
    git clone https://github.com/Bantamlak12/personal_book_library_manager.git
    ```

2. Change your directory to the cloned one
    ```
    cd personal_book_library_manager
    ```

3. Install Go dependencies
    ```
    go mod download
    ```

4. Build the application
    ```
    go build -o bin/book-library cmd/api/main.go

## Runing the application
1. Start the server

    ```
    ./bin/api
    ```

2. The API will be available at `http://localhost:8080`
