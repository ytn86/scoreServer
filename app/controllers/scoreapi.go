package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"github.com/ytn86/scoreServer/app/models"
)

type APIv1ScoreController struct {
	App
}

func (c APIv1ScoreController) GetUsers() revel.Result {

	jsonResponse := make(map[string]interface{})

	/*
		c.Begin()
		users := models.GetTop100Users(c.Tx)
		c.Commit()
	*/

	var ranks []models.UserScore

	err := cache.Get("ranks", &ranks)

	if err != nil {

		log.Println(err)

		c.Begin()
		ranks = models.GetScores(c.Tx)
		jsonScore, _ := json.Marshal(ranks)
		models.AddScoreBoard(c.Tx, string(jsonScore))
		c.Commit()

		/*
			str,_ := json.Marshal(ranks)
			log.Println(string(str))
		*/
		go cache.Set("ranks", ranks, 1*time.Minute)
	}

	if ranks == nil {
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	if len(ranks) == 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "user not found"

		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"
	jsonResponse["data"] = ranks

	return c.RenderJSON(jsonResponse)
}

func (c APIv1ScoreController) GetTop100Users() revel.Result {

	jsonResponse := make(map[string]interface{})

	/*
		c.Begin()
		users := models.GetTop100Users(c.Tx)
		c.Commit()
	*/

	var ranks []models.UserScore

	err := cache.Get("ranks", &ranks)

	if err != nil {

		log.Println(err)

		c.Begin()
		ranks = models.GetScoresLimit100(c.Tx)
		jsonScore, _ := json.Marshal(ranks)
		models.AddScoreBoard(c.Tx, string(jsonScore))
		c.Commit()

		/*
			str,_ := json.Marshal(ranks)
			log.Println(string(str))
		*/
		go cache.Set("ranks", ranks, 1*time.Minute)
	}

	if ranks == nil {
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	if len(ranks) == 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "user not found"

		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"
	jsonResponse["data"] = ranks

	return c.RenderJSON(jsonResponse)
}
