package controllers

import (
	"github.com/revel/revel"

)

type TaskController struct {
	App
}

func (c TaskController) ViewTasks() revel.Result {
	
	return c.Render()
}

