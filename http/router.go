package router

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"milpost.ch/api"
	"milpost.ch/auth"
)

var allowedOrigins = handlers.AllowedOrigins([]string{
	"http://localhost:4200"})

// HandleHTTP Function
func HandleHTTP(port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login", auth.Use(auth.Login, auth.BasicAuth())).Methods("OPTIONS", "POST")
	router.HandleFunc("/post", api.GetPost).Methods("OPTIONS", "GET")

	router.HandleFunc("/post/{id}", api.CreatePostEntry)
	router.HandleFunc("/post/{id}", api.DeletePostEntry).Methods("DELETE")

	router.HandleFunc("/", api.GetLandingPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins)(router)))
}
