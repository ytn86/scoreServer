package controllers

import (
	"github.com/revel/revel"
)

type Admin_GenreController struct {
	App
}

func (c Admin_GenreController) ViewGenres() revel.Result {

	//check admin

	if c.Session["isAdmin"] != "1" {
		return c.Forbidden("You want to be an admin?")
	}
	
	
	return c.Render()
}
