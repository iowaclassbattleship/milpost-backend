package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONErrorModel struct {
	Error string `json:"error"`
}

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IsError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func JSONError(w http.ResponseWriter, err JSONErrorModel, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
