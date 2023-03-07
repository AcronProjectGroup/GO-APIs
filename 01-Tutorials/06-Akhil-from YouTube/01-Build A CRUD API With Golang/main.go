package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1",
		Isbn: "438227",
		Title: "Movie One",
		Director: &Director{
			Firstname: "Sina",
			Lastname: "Lale Bakhsh",
		},
	})


	movies = append(movies, Movie{
		ID: "2",
		Isbn: "45455",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Mina",
			Lastname: "Sadr",
		},
	})


	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Startinf server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
