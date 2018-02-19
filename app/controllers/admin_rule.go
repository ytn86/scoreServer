package controllers

import (
	"github.com/revel/revel"
)

type Admin_RuleController struct {
	App
}

func (c Admin_RuleController) ViewRule() revel.Result {

	//check admin
	if c.Session["isAdmin"] != "1" {

		return c.Forbidden("You want to be admin?")

	}

	return c.Render()
}
