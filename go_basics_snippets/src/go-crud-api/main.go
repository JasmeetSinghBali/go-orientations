package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	// the json is the req-resp in small key value pairs in lowercase
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// slice of type Movie ac tas collection of movies DB
var movies []Movie

/*Get movies - ROUTEHandler*/
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// encoding the reponse into json and returning all movies i.e slice Movie
	json.NewEncoder(w).Encode(movies)
	return
}

/*Delete movie param=ID- ROUTEHandler*/
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// grab the params passed in r request object via mux.Vars(reqObject)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete this movie id entry, by appending all other entries before the index that we want to delete [:index] & after the index to be deleted i.e [index+1:]... just excluding this entry ID passed as params that we want to delete
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// return the remaining movies
	json.NewEncoder(w).Encode(movies)
	return
}

/*Get a single movie - ID routeHandler*/
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

/*Create movie - routeHandler*/
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// to hold the request payload for movie creation in postman
	var movie Movie
	// decode the json body in postman request
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
	return
}

/*update movie params=id- routeHandler*/
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// delete the movie from slice first
			movies = append(movies[:index], movies[index+1:]...)
			// add the new movie a/c to request payload
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	// r is now a router
	r := mux.NewRouter()

	// populating the Movie slice with dummy movies
	movies = append(movies, Movie{ID: "1", Isbn: "7440356", Title: "Movie One", Director: &Director{FirstName: "Jacob", LastName: "Mamoa"}})
	movies = append(movies, Movie{ID: "2", Isbn: "6430350", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Wick"}})

	// CRUD routes
	// getMovies is basically the handler that handles /movies route requests
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at 5000\n")
	log.Fatal(http.ListenAndServe(":5000", r))
}
