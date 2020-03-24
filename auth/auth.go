package auth

import (
	"net/http"

	"milpost.ch/errorhandler"
)

func BasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, isOk := r.BasicAuth()

		if "user" != user || "pass" != pass || isOk == false {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			errorhandler.JSONError(w, errorhandler.JSONErrorModel{"Unauthorized"}, http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hello"))
}
