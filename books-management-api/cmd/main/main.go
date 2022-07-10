package main

import (
	"log"
	"net/http"

	"github.com/Jasmeet-1998/go-orientations/books-management-api/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// create a new app router and ppas it to the RegisterBookStoreRoutes func
	appRouter := mux.NewRouter()
	routes.RegisterBookStoreRoutes(appRouter)
	http.Handle("/", appRouter)
	log.Fatal(http.ListenAndServe("localhost:8000", appRouter))
}
