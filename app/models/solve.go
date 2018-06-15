package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Solve struct {
	ID        int       `json:"-"`
	UserID    int       `json:"-"`
	TaskID    int       `json:"taskid,ormitempty"`
	SolvedAt  time.Time `json:"solved_at,ormitempty"`
	Answer    string    `json:"answer"`
	IsCorrect bool      `json:"is_correct"`
	Title     string    `sql:"-" json:"title,ormitempty"`
	Point     int       `sql:"-" json:"point,ormitempty"`
}

type UserScore struct {

	//if use User instead of Name
	//gorm says 'no such column solves.deleted_at'
	// but I do not know why.  May be instead of users.deleted_at mistakely ?

	ID           int    `json:"id,ormitempty"`
	Rank         int    `json:"rank"`
	Name         string `json:"name"`
	Comment      string `json:"comment"`
	LastSolvedAt string `json:"last_solved_at"`
	Point        int    `json:"point"`
}

func NewSolvedTask(userID int, taskID int) Solve {

	var solve Solve
	solve.UserID = userID
	solve.TaskID = taskID
	solve.SolvedAt = time.Now()

	return solve
}

func NewSubmission(userID int, taskID int, answer string, isCorrect bool) Solve {

	var solve Solve
	solve.UserID = userID
	solve.TaskID = taskID
	solve.Answer = answer
	solve.IsCorrect = isCorrect
	solve.SolvedAt = time.Now()

	return solve
}

func GetSolvedTasksByUserID(dbses *gorm.DB, userID int) []Solve {

	var solves []Solve

	err := dbses.Table("solves").
		Select([]string{"solves.task_id", "solves.solved_at", "tasks.title", "tasks.point"}).
		Where("solves.user_id = ?", userID).
		Where("solves.is_correct = ?", true).
		Joins("left join tasks on tasks.id = solves.task_id").
		Find(&solves).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	if len(solves) == 0 {
		return []Solve{}
	}

	return solves
}

func AddSolvedTask(dbsess *gorm.DB, userID int, taskID int) bool {

	solve := NewSolvedTask(userID, taskID)

	err := dbsess.Table("solves").
		Create(&solve).Error

	if err != nil {
		return false
	}

	return true

}

func AddSubmission(dbsess *gorm.DB, userID int, taskID int, answer string, isCorrect bool) bool {

	solve := NewSubmission(userID, taskID, answer, isCorrect)

	err := dbsess.Table("solves").
		Create(&solve).Error

	if err != nil {
		return false
	}

	return true

}

func IsTaskSolved(dbsess *gorm.DB, userID int, taskID int) bool {

	var solve Solve
	err := dbsess.Table("solves").
		Where("solves.user_id = ?", userID).
		Where("solves.task_id = ?", taskID).
		Where("solves.is_correct", true).
		First(&solve).Error

	if err != nil {
		return false
	}

	if solve.ID == 0 {
		return false
	}

	return true

}

func GetScores(dbsess *gorm.DB) []UserScore {

	var scores []UserScore

	err := dbsess.Table("solves").
		Select([]string{"users.name", "sum(tasks.point) as point", "max(solves.solved_at) as last_solved_at", "users.comment"}).
		Group("user_id").
		Joins("join users on users.id = solves.user_id").
		Joins("join tasks on solves.task_id = tasks.id").
		Where("tasks.is_available = ?", true).
		Where("sovles.is_correct = ?", true).
		Order("point desc, last_solved_at asc").
		Find(&scores).Error

	if err != nil {
		log.Println(err)
		return []UserScore{}
	}

	for i, _ := range scores {
		scores[i].Rank = i + 1
	}
	return scores

}

func GetScoresLimit100(dbsess *gorm.DB) []UserScore {

	var scores []UserScore

	err := dbsess.Table("solves").
		Select([]string{"users.id", "users.name", "sum(tasks.point) as point", "max(solves.solved_at) as last_solved_at", "users.comment"}).
		Group("user_id").
		Joins("join users on users.id = solves.user_id").
		Joins("join tasks on solves.task_id = tasks.id").
		Where("tasks.is_available = ?", true).
		Order("point desc, last_solved_at asc").
		Limit(100).
		Find(&scores).Error

	if err != nil {
		log.Println(err)
		return []UserScore{}
	}

	for i, _ := range scores {
		scores[i].Rank = i + 1
	}
	return scores

}

func GetUserScore(dbsess *gorm.DB, userID int) int {

	var score UserScore

	err := dbsess.Table("solves").
		Select("sum(point) as point").
		Joins("join tasks on tasks.id = solves.task_id").
		Where("solves.user_id = ?", userID).
		Find(&score).Error

	if err != nil {
		log.Println(err)
		return -1
	}

	return score.Point
}
