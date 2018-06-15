package main

import (
	"../models"
	"./db"
	
	"github.com/jinzhu/gorm"
)

var (
	dbsession *gorm.DB
)

func migrate(dbsess *gorm.DB) {

	dbsess.AutoMigrate(&models.Task{}, &models.User{}, &models.Genre{}, &models.Solve{}, &models.ScoreBoard{})
}

func main() {

	dbsess := db.InitDB("../../db/scoreServer.db")
	migrate(dbsess)
}
