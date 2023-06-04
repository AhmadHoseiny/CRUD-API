# CRUD API

This a simple CRUD API using GO. It's implemented for educational purposes. 
It supports the following operations:

- Create a new movie
- Get a movie by ID
- Get all movies
- Update a movie by ID
- Delete a movie by ID

Note: The API is not connected to a database, so the data is not persisted.

## How to run

1. Clone the repository
2. Run `go run main.go`
3. Open Postman and test the API using the following endpoints:
    - Create a new movie: `POST http://localhost:5000/movies`
    - Get a movie by ID: `GET http://localhost:5000/movies/{id}`
    - Get all movies: `GET http://localhost:5000/movies`
    - Update a movie by ID: `PUT http://localhost:5000/movies/{id}`
    - Delete a movie by ID: `DELETE http://localhost:5000/movies/{id}`

