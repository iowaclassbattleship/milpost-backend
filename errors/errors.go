package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONErrorModel struct {
	Error string `json:"error"`
}

func Fatal(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ErrorHandlerInternal(w http.ResponseWriter, err error, message string, status int) {
	if err != nil {
		JSONError(w, JSONErrorModel{Error: message}, status)
	}
}

func JSONError(w http.ResponseWriter, err JSONErrorModel, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
