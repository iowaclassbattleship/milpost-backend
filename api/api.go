package api

import (
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	"milpost.ch/db"
	errorhandler "milpost.ch/errorhandler"
	"milpost.ch/model"
)

type envelope struct {
	Result result `json:"result"`
}

type result struct {
	Meta meta         `json:"meta"`
	Data []model.Post `json:"post"`
}

type meta struct {
	StatusCode int       `json:"statusCode"`
	Time       time.Time `json:"time"`
}

func buildResponse() []byte {
	var response []model.Post
	for i := 0; i < 5; i++ {
		response = append(response, model.Post{
			"Charlie",
			"Ambos",
			"General",
			"Fishman",
			2,
		})
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
	var p model.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	errorhandler.Fatal(err)

	db.InsertPost(p)
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

func buildResponseEnvelope(response []model.Post) envelope {
	return envelope{result{meta{200, time.Now()}, response}}
}
