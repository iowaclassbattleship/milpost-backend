package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"milpost.ch/errorhandler"
)

type tokenResult struct {
	Token string `json:"token"`
}

type claims struct {
	Username string `json:"user"`
	jwt.StandardClaims
}

const (
	KEY = "123"
)

func BasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, isOk := r.BasicAuth()

		if "user" != user || "pass" != pass || isOk == false {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			errorhandler.JSONError(w, errorhandler.JSONErrorModel{Error: errorhandler.Unauthorized}, http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func JWTAuth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("Authorization")
		a = strings.Replace(a, "Bearer ", "", 1)

		cl := &claims{}
		tkn, err := jwt.ParseWithClaims(a, cl, func(token *jwt.Token) (interface{}, error) {
			return []byte("1234"), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			fmt.Print("invalid token")
		}

		fmt.Print(cl.Username)

		h.ServeHTTP(w, r)
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(45 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "yuppie",
		"exp":  expiration,
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(KEY))
	if errorhandler.IsError(err) == true {
		errorhandler.JSONError(w, errorhandler.JSONErrorModel{Error: errorhandler.TokenGenerationFailed}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResult{tokenString})
}
