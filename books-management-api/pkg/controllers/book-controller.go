package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jasmeet-1998/go-orientations/books-management-api/pkg/models"
	"github.com/Jasmeet-1998/go-orientations/books-management-api/pkg/utils"
	"github.com/gorilla/mux"
)

/* NewBook of type Book struct defined in pkg/models*/
var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	allBooks := models.GetAllBooks()
	// Marshal function to convert the data from DB into json i.e serializing the data from db
	res, _ := json.Marshal(allBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200 status code
	// return the json version of the books to the user
	w.Write(res)
	return
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	// string---> int explicit conversion
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing string [bookId] to int")
	}
	// not caring about the db instance
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// store the reff for the Book struct from models package
	CreateBook := &models.Book{}
	// parse the request and store the deserialize version in CreateBook var
	utils.ParseBody(r, CreateBook)
	// call CreateBook method to create new book with deserialize request payload
	newBook := CreateBook.CreateBook()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing [bookID]")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	// parse & store the new data for the book in updateBook variable
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing [bookId]")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	// save the new book as with updated book payload from user
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}
