package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleHTTP Function
func HandleHTTP(port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login", Login).Methods("OPTIONS", "POST")
	router.HandleFunc("/post", GetPost).Methods("OPTIONS", "GET")

	router.HandleFunc("/post/{id}", CreatePostEntry).Methods("POST")
	router.HandleFunc("/post/{id}", DeletePostEntry).Methods("DELETE")

	router.HandleFunc("/", GetLandingPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
