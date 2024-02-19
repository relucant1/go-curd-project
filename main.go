package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:id`
	Ishon    string    `json:ishon`
	Title    string    `json:title`
	Director *Director `json:director`
}

type Director struct {
	Firstname string `json:firstname`
	Lasttname string `json:lastname`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, v := range movies {
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies = append(movies, Movie{ID: "3", Ishon: "455343", Title: "Move three", Director: &Director{Firstname: "Ben", Lasttname: "Simens"}})
	json.NewEncoder(w).Encode(movies)
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Ishon: "4343", Title: "Move one", Director: &Director{Firstname: "John", Lasttname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Ishon: "3432", Title: "Move two", Director: &Director{Firstname: "Steve", Lasttname: "Smith"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/createMovie", createMovie).Methods("GET")
	//r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("hello world go server")
	log.Fatal(http.ListenAndServe(":8080", r))
}
