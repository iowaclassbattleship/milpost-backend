package api

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"milpost.ch/auth"
)

var allowedOrigins = handlers.AllowedOrigins([]string{
	"http://localhost:4200"})

// HandleHTTP Function
func HandleHTTP(port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login", Login).Methods("OPTIONS", "POST")
	router.HandleFunc("/post", GetPost).Methods("OPTIONS", "GET")

	router.HandleFunc("/post/{id}", auth.Use(CreatePostEntry, auth.BasicAuth))
	router.HandleFunc("/post/{id}", DeletePostEntry).Methods("DELETE")

	router.HandleFunc("/", GetLandingPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins)(router)))
}
