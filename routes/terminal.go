package routes

import (
	"GO-TEST-MKP/controllers"
	"GO-TEST-MKP/middleware"

	"github.com/gorilla/mux"
)

func TerminalRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()
	router.Use(middleware.Auth)

	router.HandleFunc("/terminal", controllers.Terminal).Methods("POST")
}