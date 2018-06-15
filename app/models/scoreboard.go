package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type ScoreBoard struct {
	ID         int
	Scoreboard string `sql:"type:json"`
	CreatedAt  time.Time
}

func newScoreBoard(sb string) *ScoreBoard {

	return &ScoreBoard{
		Scoreboard: sb,
		CreatedAt:  time.Now(),
	}

}

func AddScoreBoard(dbsess *gorm.DB, sb string) bool {

	err := dbsess.Table("score_boards").Create(newScoreBoard(sb)).Error

	if err != nil {
		log.Println(err)
		return false
	}

	return true

}
