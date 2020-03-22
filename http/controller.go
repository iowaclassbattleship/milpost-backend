package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HandleHTTP Function
func HandleHTTP(port int, uri string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/post/all", PostHandler)
	router.HandleFunc("/", WelcomeHandler)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
