package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       int
	Data      BookTicket
	TimeStamp string
	Hash      string
	PrevHash  string
}

// 📝 used by newBook /new POST
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

// 📝 used by writeBlock / POST
type BookTicket struct {
	BookID     string `json:"book_id"`
	User       string `json:"user"`
	BoughtDate string `json:"bought_date"`
	isGenesis  bool   `json:"is_genesis"`
}

type Blockchain struct {
	// slice of type of Block
	blocks []*Block
}

var Blockchain *Blockchain

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET") // get entire blockchain
	r.HandleFunc("/", writeBlock).Methods("POST")   // write on the block @returns BookID that is getting purchased and is passed to BookTicket
	r.HandleFunc("/new", newBook).Methods("POST")   // to create new book

	log.Println("Listening on port 5000")

	log.Fatal(http.ListenAndServe(":5000", r))
}
