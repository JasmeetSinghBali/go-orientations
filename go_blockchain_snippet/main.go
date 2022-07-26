package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET") // get entire blockchain
	r.HandleFunc("/", writeBlock).Methods("POST")   // write on the block
	r.HandleFunc("/new", newBook).Methods("POST")   // to create new book

	log.Println("Listening on port 5000")

	log.Fatal(http.ListenAndServe(":5000", r))
}
