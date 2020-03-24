package errorhandler

import (
	"encoding/json"
	"net/http"
)

type JSONErrorModel struct {
	Error string `json:"error"`
}

// ErrorHandler checks, if an error occured
func ErrorHandler(err error) {
	if err != nil {
		panic(err)
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
