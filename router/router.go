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
	"*"})

var allowedHeaders = handlers.AllowedHeaders([]string{
	"Access-Control-Allow-Origin", "X-Requested-With", "Content-Type", "Authorization"})

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
	router.HandleFunc("/post", auth.JWTAuth(api.CreatePostEntry)).Methods("POST")

	router.HandleFunc("/post/{id}", auth.JWTAuth(api.DeletePostEntry)).Methods("DELETE")

	router.HandleFunc("/", api.GetLandingPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(router)))
}
