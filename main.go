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

// Constants
const ContentType = "Content-Type"
const Json = "application/json"
const PORT = ":8000"

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Movie Slice. Our in memory DB
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, Json)
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, Json)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, Json)
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				fmt.Printf("Error %v\n", err)
			}
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, Json)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, Json)
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	r := mux.NewRouter()
	// Initial data
	movies = append(movies, Movie{ID: "1", ISBN: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Cena"}})
	movies = append(movies, Movie{ID: "2", ISBN: "456444", Title: "Movie Two", Director: &Director{FirstName: "Jane", LastName: "Cena"}})

	// Setup Routes
	r.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	r.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)
	r.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)

	fmt.Printf("Starting server at %s\n", PORT)
	// Start the server
	log.Fatal(http.ListenAndServe(PORT, r))
}
