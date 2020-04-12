package api

import (
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	errorhandler "milpost.ch/errorhandler"
)

type envelope struct {
	Result result `json:"result"`
}

type result struct {
	Meta meta   `json:"meta"`
	Data []post `json:"post"`
}

type meta struct {
	StatusCode int       `json:"statusCode"`
	Time       time.Time `json:"time"`
}

type post struct {
	ID        int    `json:"id"`
	Company   string `json:"company"`
	Section   string `json:"section"`
	Grade     string `json:"grade"`
	Name      string `json:"name"`
	ItemType  int    `json:"itemType"`
	Timestamp string `json:"timeStamp"`
}

func buildResponse() []byte {
	var response []post
	for i := 0; i < 5; i++ {
		response = append(response, post{
			1,
			"Charlie",
			"Ambos",
			"General",
			"Fishman",
			2,
			"2019-04-03"})
	}

	js, err := json.Marshal(buildResponseEnvelope(response))
	errorhandler.ErrorHandler(err)

	return js
}

// GetPost returns all entries
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(buildResponse())
}

// CreatePostEntry creates a new entry in the database and returns an id
func CreatePostEntry(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hfhfhfh"))
}

// DeletePostEntry deletes an entry in the database
func DeletePostEntry(w http.ResponseWriter, r *http.Request) {
	http.StatusText(404)
}

// GetLandingPage returns a html file welcoming the user
func GetLandingPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles("./templates/index.html")
	errorhandler.ErrorHandler(err)

	t.Execute(w, nil)
}

func buildResponseEnvelope(response []post) envelope {
	return envelope{result{meta{200, time.Now()}, response}}
}
