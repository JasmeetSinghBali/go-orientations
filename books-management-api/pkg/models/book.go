package models

import (
	"github.com/Jasmeet-1998/go-orientations/books-management-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name`
	Author      string `json:"author"`
	Publication string `json:"publication`
}

// to initialize the db connection
func init() {
	config.Connect()
	db = config.GetDB()
	// migrate the Book schema/struct to the DB
	db.AutoMigrate(&Book{})
}
