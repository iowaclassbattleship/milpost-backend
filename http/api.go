package api

import (
	"encoding/json"
	"net/http"
	"text/template"
)

type post struct {
	ID           int
	Company      string
	Section      string
	Rank         string
	Name         string
	DeliveryType int
	Timestamp    string
}

func buildResponse() []byte {
	response := post{
		1,
		"Charlie",
		"Ambos",
		"General",
		"Fishman",
		2,
		"2019-04-03"}

	js, err := json.Marshal(response)
	ErrorHandler(err)

	return js
}

// GetPost returns all entries
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(buildResponse())
}

// CreatePostEntry creates a new entry in the database and returns an id
func CreatePostEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(buildResponse())
}

// DeletePostEntry deletes an entry in the database
func DeletePostEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(buildResponse())
}

// GetLandingPage returns a html file welcoming the user
func GetLandingPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles("./templates/index.html")
	ErrorHandler(err)

	t.Execute(w, nil)
}
