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
	router.HandleFunc("GET /product/{productID}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get product by ID" + r.PathValue("productID")))
	})
	// 💡 catch all request with this handler that has no METHOD specified
	router.HandleFunc("/product/{productID}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Catch all request"))
	})

	// 📌 sub routing
	// grouping all /api/v1 route to go to router
	v1 := http.NewServeMux()
	// 💡 http.StripPrefix("/api/v1", router) strips /api/v1 from the URL path, transforming /api/v1/products/123 into /products/123.
	// 💡 The transformed path /products/123 is then handled by the main router. If the main router has a handler registered for /products/{productID},
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	// order matters if auth middleware is called first then the control does not even goes to the request logger middleware
	middlewarechain := MiddlewareChain(
		RequestLoggerMiddleware,
		RequireAuthMiddleware,
	)
	server := http.Server{
		Addr: s.addr,
		//  note here middlewarechain is inferenced as RequireAuthMiddleware(RequestLoggerMiddleware(router))
		Handler: middlewarechain(router), // 📌 using middlewares // change to v1 for subrouter
	}
	log.Printf("Server has started started %s", s.addr)
	return server.ListenAndServe()
}

// simple request logger middleware for incoming request wraps the router
func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method: %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

// middleware to authenticate authorization token or api key for incoming req to setup protected routes
func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if the user is authenticated
		token := r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return // return from the middleware
		}
		// if authenticated then call next
		next.ServeHTTP(w, r)

	}
}

type Middleware func(http.Handler) http.HandlerFunc

// A variadic parameter ... is a parameter that can accept zero or more arguments.
func MiddlewareChain(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		// applied in reverse order the last middleware in the chain will be a wrapper for first one i.e AuthLogger(RequsetLogger(router)) so that the inner one executes first
		for i := len(middleware) - 1; i >= 0; i-- {
			log.Println(middleware[i])
			// For each middleware function, it wraps the current next handler with the middleware, creating a new handler.
			// The newly created handler is then assigned back to next, so it becomes the handler to be wrapped by the next middleware in the iteration.
			next = middleware[i](next)
		}
		return next.ServeHTTP
	}
}
