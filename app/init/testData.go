package main

import (
	_ "fmt"

	"../models"
	"./db"
)

func main() {

	//dbsession := db.InitDB("../db/scoreServer.db")
	dbsess := db.InitDB("../../db/scoreServer.db")

	genre1 := &models.Genre{Name: "pwn", Color: "#ffc000"}
	genre2 := &models.Genre{Name: "web", Color: "#7fffd4"}
	genre3 := &models.Genre{Name: "crypto", Color: "#ffd700"}
	genre4 := &models.Genre{Name: "misc", Color: "#b7b7ff"}

	dbsess.Create(genre1)
	dbsess.Create(genre2)
	dbsess.Create(genre3)
	dbsess.Create(genre4)

	var pwn models.Genre
	var web models.Genre
	var cry models.Genre
	var msc models.Genre

	dbsess.Where(&models.Genre{Name: "web"}).First(&web)
	dbsess.Where(&models.Genre{Name: "pwn"}).First(&pwn)
	dbsess.Where(&models.Genre{Name: "crypto"}).First(&cry)
	dbsess.Where(&models.Genre{Name: "misc"}).First(&msc)

	task1 := &models.Task{GenreID: web.ID, Title: "test1", Description: "test1", Flag: "FLAG{test1}", Point: 10, IsAvailable: true}
	task2 := &models.Task{GenreID: web.ID, Title: "test2", Description: "test2", Flag: "FLAG{test2}", Point: 20, IsAvailable: true}
	task3 := &models.Task{GenreID: web.ID, Title: "test3", Description: "test3", Flag: "FLAG{test3}", Point: 40, IsAvailable: true}
	task4 := &models.Task{GenreID: pwn.ID, Title: "test4", Description: "test4", Flag: "FLAG{test4}", Point: 30, IsAvailable: true}
	task5 := &models.Task{GenreID: pwn.ID, Title: "test5", Description: "test5", Flag: "FLAG{test5}", Point: 33, IsAvailable: true}
	task6 := &models.Task{GenreID: pwn.ID, Title: "test6", Description: "test6", Flag: "FLAG{test6}", Point: 12, IsAvailable: true}
	task7 := &models.Task{GenreID: cry.ID, Title: "test7", Description: "test7", Flag: "FLAG{test7}", Point: 55, IsAvailable: true}
	task8 := &models.Task{GenreID: cry.ID, Title: "test8", Description: "test8", Flag: "FLAG{test8}", Point: 91, IsAvailable: true}
	task9 := &models.Task{GenreID: pwn.ID, Title: "test9", Description: "test9", Flag: "FLAG{test9}", Point: 11, IsAvailable: true}
	taska := &models.Task{GenreID: msc.ID, Title: "testa", Description: "testa", Flag: "FLAG{testa}", Point: 44, IsAvailable: true}
	taskb := &models.Task{GenreID: msc.ID, Title: "testb", Description: "testb", Flag: "FLAG{testb}", Point: 100, IsAvailable: true}
	taskc := &models.Task{GenreID: msc.ID, Title: "testc", Description: "testc", Flag: "FLAG{testc}", Point: 200, IsAvailable: true}
	user1 := models.NewUser("hoge", "hoge", "password", true)
	user2 := models.NewUser("fuga", "fuga", "password", true)
	user3 := models.NewUser("mayuge", "mayuge", "password", true)
	user4 := models.NewUser("hoge1", "hoge1", "password", true)
	user5 := models.NewUser("hoge2", "hoge2", "password", true)
	user6 := models.NewUser("hoge3", "hoge3", "password", true)
	user7 := models.NewUser("hoge4", "hoge4", "password", false)
	user8 := models.NewUser("hoge5", "hoge5", "password", true)
	user9 := models.NewUser("hoge6", "hoge6", "password", true)
	usera := models.NewUser("hoge7", "hoge7", "password", false)
	userb := models.NewUser("hoge8", "hoge8", "password", true)
	userc := models.NewUser("hoge9", "hoge9", "password", true)
	userd := models.NewUser("hogea", "hogea", "password", true)
	usere := models.NewUser("hogeb", "hogeb", "password", true)
	userf := models.NewUser("hogec", "hogec", "password", true)
	user10 := models.NewUser("hoged", "hoged", "password", true)


	dbsess.Create(task1)
	dbsess.Create(task2)
	dbsess.Create(task3)
	dbsess.Create(task4)
	dbsess.Create(task5)
	dbsess.Create(task6)
	dbsess.Create(task7)
	dbsess.Create(task8)
	dbsess.Create(task9)
	dbsess.Create(taska)
	dbsess.Create(taskb)
	dbsess.Create(taskc)
	dbsess.Create(user1)
	dbsess.Create(user2)
	dbsess.Create(user3)
	dbsess.Create(user4)
	dbsess.Create(user5)
	dbsess.Create(user6)
	dbsess.Create(user7)
	dbsess.Create(user8)
	dbsess.Create(user9)
	dbsess.Create(usera)
	dbsess.Create(userb)
	dbsess.Create(userc)
	dbsess.Create(userd)
	dbsess.Create(usere)
	dbsess.Create(userf)
	dbsess.Create(user10)

	models.AddSubmission(dbsess, 1, 1, "", true)
	models.AddSubmission(dbsess, 1, 2, "", true)
	models.AddSubmission(dbsess, 1, 3, "", true)
	models.AddSubmission(dbsess, 1, 4, "", true)
	models.AddSubmission(dbsess, 1, 5, "", true)
	models.AddSubmission(dbsess, 1, 6, "", true)
	models.AddSubmission(dbsess, 2, 1, "", true)
	models.AddSubmission(dbsess, 2, 3, "", true)
	models.AddSubmission(dbsess, 2, 4, "", true)
	models.AddSubmission(dbsess, 2, 5, "", true)
	models.AddSubmission(dbsess, 2, 6, "", true)
	models.AddSubmission(dbsess, 3, 1, "", true)
	models.AddSubmission(dbsess, 3, 4, "", true)
	models.AddSubmission(dbsess, 3, 7, "", true)
	models.AddSubmission(dbsess, 3, 8, "", true)
	models.AddSubmission(dbsess, 4, 1, "", true)
	models.AddSubmission(dbsess, 4, 2, "", true)
	models.AddSubmission(dbsess, 4, 3, "", true)
	models.AddSubmission(dbsess, 4, 4, "", true)
	models.AddSubmission(dbsess, 4, 5, "", true)
	models.AddSubmission(dbsess, 5, 1, "", true)
	models.AddSubmission(dbsess, 5, 2, "", true)
	models.AddSubmission(dbsess, 5, 6, "", true)
	models.AddSubmission(dbsess, 5, 7, "", true)
	models.AddSubmission(dbsess, 5, 9, "", true)
	models.AddSubmission(dbsess, 6, 5, "", true)
	models.AddSubmission(dbsess, 6, 4, "", true)
	models.AddSubmission(dbsess, 6, 8, "", true)
	models.AddSubmission(dbsess, 6, 9, "", true)
	
	
}
