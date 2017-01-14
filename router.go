package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Creates a router that handles the routes specified in the routes.go file
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(corsHandler(route.HandlerFunc))
	}

	return router
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
