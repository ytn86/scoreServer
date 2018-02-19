package controllers

import (
	//"encoding/json"
	//"io/ioutil"
	"log"
	"strconv"
	//"time"

	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"github.com/ytn86/scoreServer/app/models"
)

type APIv1TaskController struct {
	App
}

//func (c APIv1TaskController) GetAllTasksInfo(taskID int) revel.Result {
func (c APIv1TaskController) GetAllAvailableTasksInfo() revel.Result {

	jsonResponse := make(map[string]interface{})
	var tasks []models.TaskWithGenre

	userID, err := strconv.Atoi(c.Session["userid"])


	if err != nil {
		/*
		log.Println(err)
		
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"
		return c.RenderJSON(jsonResponse)
*/
		userID = 0
	}
	/*
	//if non-logined user : "availableTasks"
	//else loged in user : "availableTasks"+UserID
	err = cache.Get("availableTasks"+c.Session["userid"], &tasks)
	if err != nil {
		log.Println(err)

		c.Begin()
		tasks = models.GetAllAvailableTasks(c.Tx, userID)
		c.Commit()

		go cache.Set("availableTasks"+c.Session["userid"], tasks, 30*time.Minute)
	}
*/

	c.Begin()
	tasks = models.GetAllAvailableTasks(c.Tx, userID)
	c.Commit()
		
	if len(tasks) == 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "task not found"
		return c.RenderJSON(jsonResponse)
	}
	// check whether this user solved the problem

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"

	jsonResponse["data"] = tasks

	return c.RenderJSON(jsonResponse)
}

func (c APIv1TaskController) GetAvailableTaskInfo(taskID int) revel.Result {

	jsonResponse := make(map[string]interface{})
	data := make(map[string]interface{})

	userID, err := strconv.Atoi(c.Session["userid"])

	if err != nil {
		userID = 0
	}
	
	c.Begin()
	task := models.GetAvailableTask(c.Tx, taskID, userID)
	c.Commit()

	if task.ID == 0 {
		jsonResponse["status"] = 404
		return c.RenderJSON(jsonResponse)
	}
	// check whether this user solved the problem

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"

	data["id"] = task.ID
	data["title"] = task.Title
	data["description"] = task.Description
	data["point"] = task.Point
	data["isSolved"] = false

	log.Println(task)
	
	jsonResponse["data"] = task
	return c.RenderJSON(jsonResponse)
}

func (c APIv1TaskController) GetAvailableTasksInfoFilterByGenreID(genreID int) revel.Result {

	jsonData := make(map[string]interface{})

	c.Begin()
	tasks := models.GetAvailableTasksFilterByGenreID(c.Tx, genreID)
	c.Commit()

	if tasks == nil {
		jsonData["status"] = 404
	} else {
		jsonData["status"] = 200
		jsonData["data"] = tasks
	}

	return c.RenderJSON(jsonData)
}

//TODO : login check
func (c APIv1TaskController) SubmitFlag(taskID int) revel.Result {

	jsonResponse := make(map[string]interface{})

	field := struct {
		Flag string `json:"flag"`
	}{}

	userID, err := strconv.Atoi(c.Session["userid"])

	if c.Session["user"] == "" {

		jsonResponse["status"] = 401
		jsonResponse["msg"] = "login first"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	task := models.GetTaskWithFlag(c.Tx, taskID)
	c.Commit()

	if task.ID == 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "task not found"
		return c.RenderJSON(jsonResponse)
	}
	/*
	content, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	if err != nil {
		log.Println(err)
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "invalid data"
		return c.RenderJSON(jsonResponse)
	}

	err = json.Unmarshal(content, &field)
    */

	err = c.Params.BindJSON(&field)

	if err != nil {
		log.Println(err)
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "invalid data"
		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200

	c.Begin()
	suc := models.IsTaskSolved(c.Tx, userID, taskID)
	c.Commit()

	if suc == true {
		jsonResponse["status"] = 200
		jsonResponse["msg"] = "already solved"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	suc = models.AddSubmission(c.Tx, userID, taskID, field.Flag, field.Flag == task.Flag)
	c.Commit()

	if suc == false {
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	log.Println(field.Flag)
	log.Println(task.Flag)
	

	if field.Flag != task.Flag {
		jsonResponse["correct"] = false
		jsonResponse["msg"] = "incorrect flag!"
		return c.RenderJSON(jsonResponse)
	}
	/*
		c.Begin()
		userID := models.GetUserIDByName(c.Tx, c.Session["user"])
		c.Commit()
	*/

	/*
		if userID < 0 {
			jsonResponse["status"] = 401
			jsonResponse["msg"] = "login first"

			return c.RenderJSON(jsonResponse)
		}

		if userID == -1 {
			jsonResponse["status"] = 500
			jsonResponse["msg"] = "server error"

			return c.RenderJSON(jsonResponse)
		}
	*/

	/*
		c.Begin()
		suc = models.AddSolvedTask(c.Tx, userID, taskID)
		c.Commit()
	*/

	// Delete score cache
	// Created by scoreapi.go
	go cache.Delete("ranks")

	jsonResponse["correct"] = true
	jsonResponse["msg"] = "congrats!"

	return c.RenderJSON(jsonResponse)

}
