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

	"milpost.ch/errors"
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
		fmt.Print(user, pass)

		if "user" != user || "pass" != pass || isOk == false {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			errors.JSONError(w, errors.JSONErrorModel{Error: errors.Unauthorized}, http.StatusUnauthorized)
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
		_, err := jwt.ParseWithClaims(a, cl, func(token *jwt.Token) (interface{}, error) {
			return getPublicKey(), nil
		})
		errors.ErrorHandlerInternal(w, err, errors.Unauthorized, http.StatusUnauthorized)
		h.ServeHTTP(w, r)
	})
}

func GetJWTRS256(w http.ResponseWriter, r *http.Request) {
	user, _, _ := r.BasicAuth()
	cl := claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, cl)
	ss, err := token.SignedString(getPrivateKey())
	errors.ErrorHandlerInternal(w, err, errors.TokenGenerationFailed, http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResult{ss})
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyBytes, err := ioutil.ReadFile("auth/keys/milpost.pem")
	errors.Fatal(err)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	errors.Fatal(err)

	return privateKey
}

func getPublicKey() *rsa.PublicKey {
	publicKeyBytes, err := ioutil.ReadFile("auth/keys/milpost.pub.pem")
	errors.Fatal(err)

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	errors.Fatal(err)

	return publicKey
}
