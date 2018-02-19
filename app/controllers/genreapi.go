package controllers

import (
	"github.com/revel/revel"
	"github.com/ytn86/scoreServer/app/models"
)

type APIv1GenreController struct {
	App
}

func (c APIv1GenreController) GetAllGenresInfo() revel.Result {

	jsonData := make(map[string]interface{})

	c.Begin()
	genres := models.GetAllGenresOrderByIDAsc(c.Tx)
	c.Commit()

	jsonData["data"] = genres

	if genres == nil {
		jsonData["status"] = 404
	} else {
		jsonData["status"] = 200
	}

	return c.RenderJSON(jsonData)
}
