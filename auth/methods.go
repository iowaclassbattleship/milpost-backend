package auth

import (
	"net/http"

	"milpost.ch/errorhandler"
)

func Method(method string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			h(w, r)
			return
		}
		errorhandler.JSONError(w, errorhandler.JSONErrorModel{Error: "Only " + method + " allowed"}, http.StatusMethodNotAllowed)
	}
}
