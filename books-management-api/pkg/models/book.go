package models

import (
	"github.com/Jasmeet-1998/go-orientations/books-management-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name`
	Author      string `json:"author"`
	Publication string `json:"publication`
}

// to initialize the db connection
func init() {
	config.ConnectToDB()
	db = config.GetDB()
	// migrate the Book schema/struct to the DB
	db.AutoMigrate(&Book{})
}

/*Services for models to interact with database*/

// CREATE new book
func (newbook *Book) CreateBook() *Book {
	// sql query for creating new book handled by gorm by default no need to write queries
	db.NewRecord(newbook)
	// use newbook value to create new book entry
	db.Create(&newbook)
	return newbook
}

// GET all books
func GetAllBooks() []Book {
	var Books []Book
	// grab all the books and store in Books slice var of type Book struct
	db.Find(&Books)
	return Books
}

// GET a book by id
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	// return both the book and the db instance of gorm.db
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
