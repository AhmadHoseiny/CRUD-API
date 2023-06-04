package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	reqID := params["id"]
	for _, movie := range Movies {
		if movie.ID == reqID {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	addMovie(w, r)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	reqID := params["id"]
	for index, movie := range Movies {
		if movie.ID == reqID {
			Movies = append(Movies[:index], Movies[index+1:]...)
			addMovie(w, r)
			return
		}
	}
	return
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(len(Movies) + 1)
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	reqID := params["id"]
	for index, movie := range Movies {
		if movie.ID == reqID {
			Movies = append(Movies[:index], Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Movies)
}
