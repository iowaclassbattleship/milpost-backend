package db

import (
	"os"

	"milpost.ch/errorhandler"
	"milpost.ch/model"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Grade    string
	Name     string
	Company  string
	Section  string
	ItemType uint8
}

func CreateTable() {
	db := dbConn()
	defer db.Close()

	db.AutoMigrate(&Post{})
}

func DummyData() {
	db := dbConn()
	defer db.Close()

	post := Post{Grade: "Fartface", Name: "Fart", Company: "hdhf", Section: "hfh", ItemType: 1}

	db.NewRecord(post)

	db.Create(&post)
}

func InsertPost(post model.Post) {
	db := dbConn()
	defer db.Close()
}

func GetPost() {
	db := dbConn()
	defer db.Close()

	var post Post
	db.Find(&post)
}

func DeletePost(id int) {
	db := dbConn()
	defer db.Close()

	db.Where("id = ?", id).Delete(&Post{})
}

func dbConn() (db *gorm.DB) {
	dialect := os.Getenv("dialect")
	dbUser := os.Getenv("dbUser")
	dbPw := os.Getenv("dbPw")
	dbName := os.Getenv("dbName")
	dbPort := os.Getenv("dbPort")
	dbSub := os.Getenv("dbSub")

	conn, err := gorm.Open(dialect, dbUser+":"+dbPw+"@tcp("+dbName+":"+dbPort+")/"+dbSub+"?parseTime=true")
	errorhandler.Fatal(err)

	return conn
}
