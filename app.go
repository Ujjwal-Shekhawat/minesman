package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
	app.Router.HandleFunc("/{command}", ServerCommand).Methods("POST")
	// app.Router.Handle("/", http.FileServer(http.Dir("./asset")))
}

func (app *App) run(port string) {
	fmt.Println("Server started on port : " + port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}
