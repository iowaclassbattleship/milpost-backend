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
		_, err := jwt.ParseWithClaims(a, cl, func(token *jwt.Token) (interface{}, error) {
			return getPublicKey(), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Print(cl.Username)

		h.ServeHTTP(w, r)
	})
}

func GetJWTRS256(w http.ResponseWriter, r *http.Request) {
	cl := claims{
		"Yolo",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, cl)
	ss, err := token.SignedString(getPrivateKey())
	if errorhandler.IsError(err) == true {
		errorhandler.JSONError(w, errorhandler.JSONErrorModel{Error: errorhandler.TokenGenerationFailed}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResult{ss})
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyBytes, err := ioutil.ReadFile("auth/keys/milpost.pem")
	fmt.Print("hruhourhg")
	errorhandler.Fatal(err)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	fmt.Print("fuck")
	errorhandler.Fatal(err)

	fmt.Print(privateKey)

	return privateKey
}

func getPublicKey() *rsa.PublicKey {
	publicKeyBytes, err := ioutil.ReadFile("auth/keys/milpost.pub.pem")
	errorhandler.Fatal(err)

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	errorhandler.Fatal(err)

	return publicKey
}
