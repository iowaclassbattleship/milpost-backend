package api

import (
	"log"
	"net/http"
	"strconv"
)

// HandleHTTP Function
func HandleHTTP(port int, uri string) {
	http.HandleFunc("/post/all", ViewHandler)
	http.HandleFunc("/", WelcomeHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
