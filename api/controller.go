package api

import (
	"log"
	"net/http"
)

// HandleHTTP Function
func HandleHTTP() {
	http.HandleFunc("/post/all", ViewHandler)
	http.HandleFunc("/", WelcomeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
