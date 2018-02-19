package controllers

import (
	"github.com/revel/revel"

	"github.com/ytn86/scoreServer/app/routes"
)

type AccountController struct {
	App
}

func (c AccountController) Register() revel.Result {

	if c.Session["user"] != "" {
		return c.Redirect(routes.TaskController.ViewTasks())
	}

	return c.Render()
}

func (c AccountController) Login() revel.Result {

	if c.Session["user"] != "" {
		return c.Redirect(routes.TaskController.ViewTasks())
	}

	return c.Render()
}

func (c AccountController) Logout() revel.Result {

	for k := range c.Session {
		delete(c.Session, k)
	}

	return c.Redirect(routes.AccountController.Login())
}

func (c AccountController) User() revel.Result {

	if c.Session["user"] == "" {
		return c.Redirect(routes.AccountController.Login())
	}

	return c.Render()
}
