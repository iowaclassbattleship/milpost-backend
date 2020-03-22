package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"text/template"
)

type post struct {
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

// ViewHandler does
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(buildResponse())
}

// WelcomeHandler does
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	ErrorHandler(err)

	fmt.Print(tmpl)
}
