package api

import (
	"encoding/json"
	"net/http"
)

type post struct {
	Company      string
	Section      string
	Rank         string
	Name         string
	DeliveryType int
	Timestamp    string
}

func buildResponse() []byte {
	response := post{
		"Charlie",
		"Ambos",
		"General",
		"Fishman",
		2,
		"2019-04-03"}

	js, err := json.Marshal(response)
	ErrorHandler(err)

	return js
}

// ViewHandler does
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(buildResponse())
}
