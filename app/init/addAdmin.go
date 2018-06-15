package main

import (
	_ "fmt"

	"../models"
	"./db"
)

func main() {

	//dbsession := db.InitDB("../db/scoreServer.db")
	dbsess := db.InitDB("../../db/scoreServer.db")

	adminUser := models.NewAdminUser("admin", "mail@ytn86.net", "adminadmin", true)

	dbsess.Create(adminUser)
	
}
