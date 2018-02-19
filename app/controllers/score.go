package controllers

import (
	"github.com/revel/revel"
)

type ScoreController struct {
	App
}

func (c ScoreController) ViewTop100Users () revel.Result {

	return c.Render()
}
