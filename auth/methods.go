package auth

import (
	"net/http"

	"milpost.ch/errorhandler"
)

func Post(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		errorhandler.JSONError(w, errorhandler.JSONErrorModel{"Only POST allowed"}, http.StatusMethodNotAllowed)
	}
}

func Get(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		errorhandler.JSONError(w, errorhandler.JSONErrorModel{"Only GET allowed"}, http.StatusMethodNotAllowed)
	}
}

func Delete(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			h(w, r)
			return
		}
		errorhandler.JSONError(w, errorhandler.JSONErrorModel{"Only DELETE allowed"}, http.StatusMethodNotAllowed)
	}
}
