package controllers

import (
	"github.com/revel/revel"
)

type RuleController struct {
	App
}

func (c RuleController) GetRule() revel.Result {

	return c.Render()
}
