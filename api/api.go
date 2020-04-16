package api

import (
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	"milpost.ch/db"
	errors "milpost.ch/errors"
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

func buildResponse(response []model.Post) []byte {
	js, _ := json.Marshal(buildResponseEnvelope(response))

	return js
}

// GetPost returns all entries
func GetPost(w http.ResponseWriter, r *http.Request) {
	post, err := db.GetPost()
	if errors.IsError(err) {
		errors.JSONError(w, errors.JSONErrorModel{Error: errors.InternalServerError}, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(buildResponse(post))
}

// CreatePostEntry creates a new entry in the database and returns an id
func CreatePostEntry(w http.ResponseWriter, r *http.Request) {
	var p model.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	errors.Fatal(err)

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
	if errors.IsError(err) == false {
		errors.JSONError(w, errors.JSONErrorModel{Error: errors.InternalServerError}, http.StatusInternalServerError)
	}

	t.Execute(w, nil)
}

func buildResponseEnvelope(response []model.Post) envelope {
	return envelope{result{meta{200, time.Now()}, response}}
}
