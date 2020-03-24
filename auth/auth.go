package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"milpost.ch/errorhandler"
)

type tokenResult struct {
	Token string `json:"token"`
}

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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "tester",
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte("123"))
	if errorhandler.IsError(err) == true {
		errorhandler.JSONError(w, errorhandler.JSONErrorModel{"TOKEN_GENERATION_FAILED"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResult{tokenString})
}
