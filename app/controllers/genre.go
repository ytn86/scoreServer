package controllers

import (
	"github.com/revel/revel"

	"github.com/ytn86/scoreServer/app/models"

)

type GenreController struct {
	App
}

func (c GenreController) GetAllGenres() revel.Result {

	var genres []models.Genre

	c.Begin()
	genres = models.GetAllGenresOrderByIDAsc(c.Tx)
	c.Commit()

	return c.Render(genres)
}

	
