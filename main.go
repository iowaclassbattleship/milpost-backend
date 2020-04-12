package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"milpost.ch/errorhandler"
	router "milpost.ch/router"
)

type environment struct {
	URI  string `json:"URI"`
	Port int    `json:"port"`
}

func main() {
	err := godotenv.Load()
	errorhandler.ErrorHandler(err)

	fmt.Println("Server listening on Port", os.Getenv("port"))

	router.HandleHTTP(os.Getenv("port"))
}
