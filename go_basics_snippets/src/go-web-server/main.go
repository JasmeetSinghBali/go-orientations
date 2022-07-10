package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	// Fprintf returns the message as response to a server request
	fmt.Fprintf(w, "hello go-web-server!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Form submitted successfully")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	// a fileserver pointing to default index.html in static folder
	fileServer := http.FileServer(http.Dir("./static"))
	// routes with handler func , where root route handle via fileserver that points to index.html by default in static folder
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server port at 5000\n")
	// listen to the server and handle any errors
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
