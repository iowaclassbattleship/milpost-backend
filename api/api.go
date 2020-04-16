package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
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
	errors.ErrorHandlerInternal(w, err, errors.InternalServerError, http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buildResponse(post))
}

// CreatePostEntry creates a new entry in the database and returns an id
func CreatePostEntry(w http.ResponseWriter, r *http.Request) {
	var p model.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	errors.ErrorHandlerInternal(w, err, errors.InternalServerError, http.StatusInternalServerError)

	insertionError := db.InsertPost(p)
	errors.ErrorHandlerInternal(w, insertionError, errors.InvalidInput, http.StatusBadRequest)
}

// DeletePostEntry deletes an entry in the database
func DeletePostEntry(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	errors.ErrorHandlerInternal(w, err, errors.InternalServerError, http.StatusInternalServerError)
	db.DeletePost(id)
}

// GetLandingPage returns a html file welcoming the user
func GetLandingPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles("./templates/index.html")
	errors.ErrorHandlerInternal(w, err, errors.InternalServerError, http.StatusInternalServerError)

	t.Execute(w, nil)
}

func buildResponseEnvelope(response []model.Post) envelope {
	return envelope{result{meta{200, time.Now()}, response}}
}
