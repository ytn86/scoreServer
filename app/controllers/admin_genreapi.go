package controllers

import (
	"log"
	//"encoding/json"
	//"io/ioutil"

	"github.com/revel/revel"
	"github.com/ytn86/scoreServer/app/models"
)

type Admin_APIv1GenreController struct {
	App
}

func (c Admin_APIv1GenreController) ModifyGenre(genreID int) revel.Result {

	var genre models.Genre
	jsonResponse := make(map[string]interface{})

	//check if admin
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

	err = json.Unmarshal(content, &genre)
*/
	err := c.Params.BindJSON(&genre)
	if err != nil {
		log.Println(err)
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "bad json"

		return c.RenderJSON(jsonResponse)
	}

	genre.ID = genreID
	c.Begin()
	suc := models.UpdateGenre(c.Tx, &genre)
	c.Commit()

	if suc == false {
		c.Response.Status = 500

		jsonResponse["status"] = 500
		jsonResponse["msg"] = "update error"

		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"

	return c.RenderJSON(jsonResponse)

}

func (c Admin_APIv1GenreController) AddGenre(genreID int) revel.Result {

	var genre models.Genre

	jsonResponse := make(map[string]interface{})

	//check admin

	if c.Session["isAdmin"] != "1" {
		jsonResponse["status"] = 403
		jsonResponse["msg"] = "You want to be an admin?"

		return c.RenderJSON(jsonResponse)
	}

	/*
	content, err := ioutil.ReadAll(c.Request.Body)
		
	if err != nil {
		log.Println(err)

		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "bad request"

		return c.RenderJSON(jsonResponse)
	}
		
	err = json.Unmarshal(content, &genre)
*/
	err := c.Params.BindJSON(&genre)
	
	if err != nil {
		log.Println(err)

		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "invalid json"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	suc := models.AddGenre(c.Tx, &genre)
	c.Commit()

	if suc == false {
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"

	return c.RenderJSON(jsonResponse)
	
}
