package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	api "milpost.ch/http"
)

type environment struct {
	URI  string `json:"URI"`
	Port int    `json:"port"`
}

func main() {
	err := godotenv.Load()
	api.ErrorHandler(err)

	fmt.Println("Server listening on Port", os.Getenv("port"))

	api.HandleHTTP(os.Getenv("port"))
}
