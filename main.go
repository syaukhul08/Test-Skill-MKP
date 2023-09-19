package main

import (
	"GO-TEST-MKP/configs"
	"GO-TEST-MKP/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
) 

func main() {
configs.ConfigDB()

r := mux.NewRouter()

router := r.PathPrefix("/api").Subrouter()

routes.AuthRoutes(router)
routes.TerminalRoutes(router)

log.Println("Server running on port 8080")
http.ListenAndServe(":8080", router)
}