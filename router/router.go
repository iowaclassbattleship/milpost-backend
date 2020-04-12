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

var allowedMethods = handlers.AllowedMethods([]string{
	"OPTIONS",
	"GET",
	"POST",
	"DELETE"})

// HandleHTTP Function
func HandleHTTP(port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login", auth.BasicAuth(auth.GetJWTRS256)).Methods("POST")

	router.HandleFunc("/post", api.GetPost).Methods("GET")

	router.HandleFunc("/post/{id}", auth.JWTAuth(api.CreatePostEntry)).Methods("POST")
	router.HandleFunc("/post/{id}", auth.JWTAuth(api.CreatePostEntry)).Methods("GET")
	router.HandleFunc("/post/{id}", auth.JWTAuth(api.DeletePostEntry)).Methods("DELETE")

	router.HandleFunc("/", api.GetLandingPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins)(router)))
}
