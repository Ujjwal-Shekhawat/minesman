package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Cors releted
var header = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
var methods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
var origins = handlers.AllowedOrigins([]string{"localhost:3000", "*"})

// export App
type App struct {
	Router *mux.Router
}

// Abstracting middleware to prevent nesting
type Middleware func(http.HandlerFunc) http.HandlerFunc

func Middlewares(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	if len(middlewares) < 1 {
		return f
	}

	wrapped_funcs := f
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped_funcs = middlewares[i](wrapped_funcs)
	}

	return wrapped_funcs
}

// Simple middleware for loggin stuff
func logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer handler.ServeHTTP(w, r)
		fmt.Println("Recived " + r.Method + " request on Path " + r.URL.Path)
	}
}

func (app *App) initRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", Login).Methods("POST", "OPTIONS")
	app.Router.HandleFunc("/console", AuthConsole).Methods("GET", "OPTIONS")
	// app.Router.Handle("/", http.FileServer(http.Dir("./asset")))
}

func (app *App) run(port string) {
	fmt.Println("Server started on port : " + port)
	log.Fatal(http.ListenAndServeTLS(":"+port,"cert.crt", "cert.key", handlers.CORS(header, methods, origins)(app.Router)))
}
