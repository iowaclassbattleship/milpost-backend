package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"milpost.ch/api"
)

type environment struct {
	URI  string `json:"URI"`
	Port int    `json:"port"`
}

func main() {
	env := readEnvironment()

	api.HandleHTTP(env.Port, env.URI)
}

func readEnvironment() environment {
	jsonFile, err := os.Open("./config/environment.json")
	api.ErrorHandler(err)

	bytes, _ := ioutil.ReadAll(jsonFile)

	var data environment

	json.Unmarshal(bytes, &data)

	return data
}
