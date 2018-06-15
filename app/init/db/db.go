package db

import (
	"log"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


func checkErr(err error, msg string) {

	log.Fatal(msg, err)
}

func InitDB(path string) *gorm.DB {

	dbsession, err := gorm.Open("sqlite3", path)
	if err != nil{
		checkErr(err, "Failed to Open")
		return nil
	}

	return dbsession

}
