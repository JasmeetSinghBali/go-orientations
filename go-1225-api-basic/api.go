package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	// 📌 declarative http methods without use of if else in http request handler
	// router.HandleFunc("GET /product/{productID}", func(w http.ResponseWriter, r *http.Request) {
	// 	// 📌 extracting path params in handlers
	// 	productID := r.PathValue("productID")
	// 	w.Write([]byte("Product ID: " + productID))
	// })
	// ❌ old way
	// func(_ w, ResponseWriter, r *http.Request){
	// 	if r.Method == 'GET'{
	// 		// handle get request
	// 		return
	// 	} else if r.Method == 'POST'{
	// 		// handle post request
	// 		return
	// 	}else {
	// 		//handle other request
	// 		return
	// 	}
	// }
	// ✅ new way [METHOD ][HOST]/[PATH] note- only 1 space allowed after METHOD else the route will be not found
	router.HandleFunc("POST /product/{productID}", func(w http.ResponseWriter, r *http.Request) {})
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}
	// 💡 catch all request with this handler that has no METHOD specified
	router.HandleFunc("/product/{productID}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Catch all request"))
	})
	log.Printf("Server has started started %s", s.addr)
	return server.ListenAndServe()
}
