package db

import (
	"os"

	"milpost.ch/errors"
	"milpost.ch/model"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
)

func CreateTable() error {
	db := dbConn()
	defer db.Close()

	db.AutoMigrate(&model.Post{})

	return nil
}

func DummyData() {
	db := dbConn()
	defer db.Close()

	post := model.Post{Grade: "Fartface", Name: "Fart", Company: "hdhf", Section: "hfh", ItemType: 1}

	db.NewRecord(post)

	db.Create(&post)
}

func InsertPost(post model.Post) {
	db := dbConn()
	defer db.Close()

	db.NewRecord(post)
	db.Create(&post)
}

func GetPost() ([]model.Post, error) {
	db := dbConn()
	defer db.Close()

	post := []model.Post{}
	db.Find(&post)

	return post, nil
}

func DeletePost(id int) error {
	db := dbConn()
	defer db.Close()

	db.Where("id = ?", id).Delete(&model.Post{})

	return nil
}

func dbConn() (db *gorm.DB) {
	dialect := os.Getenv("dialect")
	dbUser := os.Getenv("dbUser")
	dbPw := os.Getenv("dbPw")
	dbName := os.Getenv("dbName")
	dbPort := os.Getenv("dbPort")
	dbSub := os.Getenv("dbSub")

	conn, err := gorm.Open(dialect, dbUser+":"+dbPw+"@tcp("+dbName+":"+dbPort+")/"+dbSub+"?parseTime=true")
	errors.Fatal(err)

	return conn
}
