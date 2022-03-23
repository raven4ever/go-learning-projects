package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raven4ever/bookstore-management/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Println("Server available at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
