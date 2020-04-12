package auth

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func GetJWTRS256(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp": time.Now().Add(45 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(getPrivateKey())
	if errorhandler.IsError(err) == true {
		errorhandler.JSONError(w, errorhandler.JSONErrorModel{Error: errorhandler.TokenGenerationFailed}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResult{tokenString})
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyByes, err := ioutil.ReadFile("auth/keys/milpost.rsa")
	errorhandler.Fatal(err)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByes)
	errorhandler.Fatal(err)

	return privateKey
}

func getPublicKey() *rsa.PublicKey {
	publicKeyBytes, err := ioutil.ReadFile("auth/keys/milpost.rsa.pub")
	errorhandler.Fatal(err)

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	errorhandler.Fatal(err)

	return publicKey
}
