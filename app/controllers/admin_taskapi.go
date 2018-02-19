package controllers

import (
	"log"

	//"encoding/json"
	//"io/ioutil"

	"github.com/revel/revel"
	//"github.com/revel/revel/cache"
	"github.com/ytn86/scoreServer/app/models"
)

type Admin_APIv1TaskController struct {
	App
}

func (c Admin_APIv1TaskController) GetAllTasksInfo() revel.Result {

	jsonResponse := make(map[string]interface{})
	
	//check admin
	if c.Session["isAdmin"] != "1" {
		jsonResponse["status"] = 403
		jsonResponse["msg"] = "You want to be an admin?"
		
		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	tasks := models.GetAllTasks(c.Tx)
	c.Commit()

	if len(tasks) == 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "task not found"

		return c.RenderJSON(jsonResponse)

	}


	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"
	jsonResponse["data"] = tasks

	return c.RenderJSON(jsonResponse)	
	
}

func (c Admin_APIv1TaskController) GetTasksInfoFilterByGenreID(genreID int) revel.Result {

	jsonResponse := make(map[string]interface{})
	
	//check admin
	/*
		username := c.Session["user"]

		c.Begin()
		isAdmin := models.IsAdmin(c.Tx, username)
		c.Commit()

		if isAdmin == false {
			jsonResponse["status"] = 403
			jsonResponse["msg"] = "You want to be an admin?"

			return c.RenderJSON(jsonResponse)
		}
	*/
	if c.Session["isAdmin"] != "1" {
		jsonResponse["status"] = 403
		jsonResponse["msg"] = "You want to be an admin?"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	tasks := models.GetTasksWithFlagFilterByGenreID(c.Tx, genreID)
	c.Commit()

	if tasks == nil {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "task not found"

		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"
	jsonResponse["data"] = tasks

	return c.RenderJSON(jsonResponse)
}

func (c Admin_APIv1TaskController) GetTaskInfo(taskID int) revel.Result {

	jsonResponse := make(map[string]interface{})

	//check admin

	if c.Session["isAdmin"] != "1" {
		jsonResponse["status"] = 403
		jsonResponse["msg"] = "You want to be an admin?"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	task := models.GetTaskWithFlag(c.Tx, taskID)
	c.Commit()

	if task.ID == -1 {
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	if task.ID == 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "task not found"

		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"
	jsonResponse["data"] = task

	return c.RenderJSON(jsonResponse)

}

/*
ToDo: check if the entered genreid is valid
*/
func (c Admin_APIv1TaskController) ModifyTask(taskID int) revel.Result {

	var task models.Task
	jsonResponse := make(map[string]interface{})

	//check admin

	if c.Session["isAdmin"] != "1" {
		jsonResponse["status"] = 403
		jsonResponse["msg"] = "You want to be an admin?"

		return c.RenderJSON(jsonResponse)
	}
	/*
	content, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	if err != nil {
		log.Println(err)
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	err = json.Unmarshal(content, &task)
*/

	err := c.Params.BindJSON(&task)
	
	if err != nil {
		log.Println(err)
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "bad json"

		return c.RenderJSON(jsonResponse)
	}


	log.Println(task)
	
	task.ID = taskID
	c.Begin()
	suc := models.UpdateTask(c.Tx, &task)
	c.Commit()

	if suc == false {
		c.Response.Status = 500

		jsonResponse["status"] = 500
		jsonResponse["msg"] = "update error"

		return c.RenderJSON(jsonResponse)
	}

	//go cache.Delete("availableTasks")

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"

	return c.RenderJSON(jsonResponse)

}

func (c Admin_APIv1TaskController) AddTask(genreID int) revel.Result {

	var task models.Task
	jsonResponse := make(map[string]interface{})

	//check admin
	
	if c.Session["isAdmin"] != "1" {
		jsonResponse["status"] = 403
		jsonResponse["msg"] = "You want to be an admin?"

		return c.RenderJSON(jsonResponse)
	}

	/*

	content, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	if err != nil {
		log.Println(err)
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	err = json.Unmarshal(content, &task)
*/

	err := c.Params.BindJSON(&task)

	if err != nil {
		log.Println(err)
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "bad json"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	suc := models.AddTask(c.Tx, &task)
	c.Commit()

	if suc == false {
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "something wrong..."

		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"

	return c.RenderJSON(jsonResponse)
}
