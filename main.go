package main

import (
	"CRUD_API/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	handlers.Movies = append(handlers.Movies, handlers.Movie{ID: "1", ISBN: "123456", Title: "Movie One", Director: &handlers.Director{Firstname: "John", Lastname: "Doe"}})
	handlers.Movies = append(handlers.Movies, handlers.Movie{ID: "2", ISBN: "654321", Title: "Movie Two", Director: &handlers.Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":5000", r))
}
