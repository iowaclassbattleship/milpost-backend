package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	db "milpost.ch/db"
	"milpost.ch/errors"
	router "milpost.ch/router"
)

type environment struct {
	URI  string `json:"URI"`
	Port int    `json:"port"`
}

func main() {
	err := godotenv.Load()
	errors.Fatal(err)
	db.CreateTable()

	fmt.Println("Server listening on Port", os.Getenv("port"))

	router.HandleHTTP(os.Getenv("port"))
}
