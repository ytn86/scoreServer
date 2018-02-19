package controllers

import (
	"github.com/revel/revel"
)

type Admin_TaskController struct {
	App
}

func (c Admin_TaskController) ViewTasks() revel.Result {

	//check admin

	if c.Session["isAdmin"] != "1" {
		return c.Forbidden("You want to be an admin?")
	}
	
	
	return c.Render()
}

func (c Admin_TaskController) AddTask() revel.Result {
	
	//check admin
	/*
	username := c.Session["user"]
	c.Begin()
	isAdmin := models.IsAdmin(c.Tx, username)
	c.Commit()
	
	if isAdmin == false {
		return c.Forbidden("You want to be an admin?")
	}
*/
	if c.Session["isAdmin"] != "1" {
		return c.Forbidden("You want to be an admin?")
	}
	
	return c.Render()
}

func (c Admin_TaskController) ModifyTask() revel.Result {

	//check admin
	/*
	username := c.Session["user"]
	c.Begin()
	isAdmin := models.IsAdmin(c.Tx, username)
	c.Commit()
	
	if isAdmin == false {
		return c.Forbidden("You want to be an admin?")
	}
*/
	if c.Session["isAdmin"] != "1" {
		return c.Forbidden("You want to be an admin?")
	}
	
	return c.Render()
}

func (c Admin_TaskController) DeleteTask() revel.Result {

	//check admin
	/*
	username := c.Session["user"]
	c.Begin()
	isAdmin := models.IsAdmin(c.Tx, username)
	c.Commit()
	
	if isAdmin == false {
		return c.Forbidden("You want to be an admin?")
	}
*/
	if c.Session["isAdmin"] != "1" {
		return c.Forbidden("You want to be an admin?")
	}
	
	return c.Render()
}
