package router

import (
	"net/http"

	"milpost.ch/api"
	"milpost.ch/auth"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var headersOk = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

// var methodsOk = handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST", "DELETE"})

// HandleHTTP Function
func HandleHTTP(port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login", auth.BasicAuth(auth.GetJWTRS256)).Methods("POST")

	router.HandleFunc("/post", api.GetPost).Methods("GET")
	router.HandleFunc("/post", auth.JWTAuth(api.CreatePostEntry)).Methods("POST")

	router.HandleFunc("/post/{id}", auth.JWTAuth(api.DeletePostEntry)).Methods("DELETE")

	router.HandleFunc("/", api.GetLandingPage).Methods("GET")

	http.ListenAndServe(":"+port, handlers.CORS(headersOk)(router))
}
