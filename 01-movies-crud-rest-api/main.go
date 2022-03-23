package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raven4ever/movies-crud-rest-api/pkg/data"
	"github.com/raven4ever/movies-crud-rest-api/pkg/model"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Println("Server available at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range data.Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie model.Movie

	_ = json.NewDecoder(r.Body).Decode(&newMovie)

	newMovie.ID = strconv.Itoa(rand.Intn(100000))

	data.Movies = append(data.Movies, newMovie)

	json.NewEncoder(w).Encode(newMovie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range data.Movies {
		if item.ID == params["id"] {
			data.Movies = append(data.Movies[:index], data.Movies[index+1:]...)

			var updatedMovie model.Movie

			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

			updatedMovie.ID = params["id"]

			data.Movies = append(data.Movies, updatedMovie)

			json.NewEncoder(w).Encode(updatedMovie)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range data.Movies {
		if item.ID == params["id"] {
			data.Movies = append(data.Movies[:index], data.Movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(data.Movies)
}
