package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectToDB() {
	// restapi is table name, jasmeet is username and Jasmeet12345 is password
	d, err := gorm.Open("mysql", "jasmeet:Jasmeet12345/restapi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// this GetDB can be called on other files so that other files can connect to the db instance
func GetDB() *gorm.DB {
	return db
}
