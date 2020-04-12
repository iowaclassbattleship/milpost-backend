package db

import (
	"database/sql"

	"milpost.ch/errorhandler"
)

func GetPost() {
	db, err := sql.Open("mysql", "")
	errorhandler.Fatal(err)

}
