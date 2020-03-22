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

// PostHandler does
func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(buildResponse())
}

// WelcomeHandler does
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles("./templates/index.html")
	ErrorHandler(err)

	t.Execute(w, nil)
}
