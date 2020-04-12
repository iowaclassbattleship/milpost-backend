package db

import (
	"database/sql"
	"fmt"
	"os"

	"milpost.ch/errorhandler"

	_ "github.com/go-sql-driver/mysql"
)

func CreateTable() {
	db := dbConn()

	_, err := db.Query(`CREATE TABLE IF NOT EXISTS post (
		post_id INT NOT NULL AUTO_INCREMENT,
		grade varchar(32),
		name varchar(128) not null,
		company varchar(32),
		section varchar(32),
		itemType BIT(1) not null,
		timestamp TIMESTAMP not null,
		CONSTRAINT post_pk PRIMARY KEY (post_id)
		)`)
	errorhandler.Fatal(err)
}

func DummyData() {
	db := dbConn()

	_, err := db.Query(`INSERT INTO post (grade, name, company, section, itemType) VALUES(
		("Lieutenant",
		"Muerner",
		"Fickschnitzel",
		"fart",
		1)
	)`)
	errorhandler.Fatal(err)
}

func Select() {
	db := dbConn()

	res, err := db.Query(`SELECT * FROM POST`)
	errorhandler.Fatal(err)

	fmt.Print(res)
}

func GetPost() {
	db := dbConn()
	fmt.Print(db)

	// res, err := db.Query("SELECT * FROM post")

}

func CreatePost() {

}

func dbConn() (db *sql.DB) {
	dialect := os.Getenv("dialect")
	dbUser := os.Getenv("dbUser")
	dbPw := os.Getenv("dbPw")
	dbName := os.Getenv("dbName")
	dbPort := os.Getenv("dbPort")
	dbSub := os.Getenv("dbSub")

	conn, err := sql.Open(dialect, dbUser+":"+dbPw+"@tcp("+dbName+":"+dbPort+")/"+dbSub)
	errorhandler.Fatal(err)

	return conn
}
