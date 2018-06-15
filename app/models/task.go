package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Task struct {
	ID          int    `gorm:"primary_key" json:"id"`
	GenreID     int    `sql:"not null" json:"genreid,omitempty"`
	Title       string `sql:"unique;not null" json:"title"`
	Description string `json:"description,omitempty"`
	Flag        string `sql:"unique;not null" json:"flag,omitempty"`
	Point       int    `sql:"not null" json:"point""`
	IsAvailable bool   `sql:"not null" json:"is_available,omitempty"`
	GenreName   string `sql:"-" json:"genre,omitempty"`
	//CreatedAt   *time.Time  `sql:"not null" json:"created_at"`
}

type TaskWithGenre struct {
	Task
	Name     string `json:"genre,omitempty"`
	Color    string `json:"color,omitempty"`
	SolvedBy int    `json:"solved_by,omitempty"`
	IsSolved bool   `json:"is_solved"`
}

func newTask(genreID int, description string, flag string, isAvailable bool) Task {

	return Task{
		GenreID:     genreID,
		Description: description,
		Flag:        flag,
		IsAvailable: isAvailable,
	}
}

func GetTaskWithFlag(dbsess *gorm.DB, taskID int) Task {

	var task Task

	err := dbsess.Table("tasks").Where("tasks.id = ?", taskID).Find(&task).Error

	if err != nil {
		log.Println(err)
		task.ID = -1

		return task
	}

	return task
}

/*
func GetAllTasksWithFlag(dbsess *gorm.DB) []Task {

	var tasks []Task
	err := dbsess.Table("tasks").Joins("join genres on genres.id = tasks.genre_id").Find(&tasks).Error

	if err != nil {
		log.Println(err)
		return []Task{}
	}

	return tasks
task}
*/

func GetAllAvailableTasks(dbsess *gorm.DB, userID int) []TaskWithGenre {
	//func GetAllTasks(dbsess *gorm.DB) []Task {
	var tasks []TaskWithGenre

	err := dbsess.Debug().Table("tasks").Select([]string{"tasks.id", "tasks.title", "tasks.point", "genres.name", "genres.color", "count(solves.task_id) as solved_by", "is_solved.is_correct as is_solved"}).
		Joins("join genres on genres.id = tasks.genre_id").
		Joins("join solves on solves.task_id = tasks.id").
		Joins("left join solves is_solved on is_solved.task_id = tasks.id AND is_solved.user_id = ?", userID).
		Where("tasks.is_available = ?", true).
		Where("solves.is_correct = ?", true).
		Group("solves.task_id").Find(&tasks).Error

	if err != nil {
		log.Println(err)
		return []TaskWithGenre{}
		//return []Task{}
	}

	return tasks
}

func GetAllTasks(dbsess *gorm.DB) []TaskWithGenre {
	//func GetAllTasks(dbsess *gorm.DB) []Task {
	var tasks []TaskWithGenre

	err := dbsess.Table("tasks").Select([]string{"tasks.id", "tasks.title", "tasks.point", "tasks.is_available", "genres.name", "genres.color", "count(solves.task_id) as solved_by"}).
		Joins("join genres on genres.id = tasks.genre_id").
		Joins("join solves on solves.task_id = tasks.id").
		Where("solves.is_correct = ?", true).
		Group("solves.task_id").Find(&tasks).Error

	if err != nil {
		log.Println(err)
		return []TaskWithGenre{}
		//return []Task{}
	}

	return tasks
}

func GetTasksWithFlagFilterByGenreID(dbsess *gorm.DB, genreID int) []Task {

	var tasks []Task
	err := dbsess.Table("tasks").Where("genre_id = ?", genreID).Find(&tasks).Error

	if err != nil {
		log.Println(err)
		return []Task{}
	}

	return tasks
}

/*
func GetAllTasksOrderByGenreAsc(dbsess *gorm.DB) []Task {

	var tasks []Task
	err := dbsess.Table("tasks").Order("genre_id asc").Find(&tasks).Error

	if err != nil {
		log.Println(err)
		return []Task{}
	}

	return tasks
}
*/

func GetAvailableTasksFilterByGenreID(dbsess *gorm.DB, genreID int) []Task {

	var tasks []Task
	//dbsess.Table("tasks").Select("id", "genre_id", "title", "description", "point").Where("genre_id = ?", genreID).Find(&tasks)
	err := dbsess.Table("tasks").Select([]string{"id", "title", "description", "point"}).
		Where("genre_id = ?", genreID).Where("tasks.is_available = ?", true).
		Find(&tasks).Error

	if err != nil {
		log.Println(err)
		return []Task{}
	}

	return tasks
}

func AddTask(dbsess *gorm.DB, task *Task) bool {

	err := dbsess.Create(&task).Error

	if err != nil {
		log.Println(err)
		return false
	}

	err = dbsess.Save(&task).Error
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

func GetAvailableTask(dbsess *gorm.DB, taskid int, userID int) TaskWithGenre {

	var task TaskWithGenre
	err := dbsess.Debug().Table("tasks").Select([]string{"tasks.id", "tasks.title", "tasks.description", "tasks.point", "solves.is_correct as is_solved"}).
		Joins("left join solves on solves.task_id = tasks.id AND solves.user_id = ?", userID).
		Where("tasks.id = ?", taskid).
		Where("tasks.is_available = ?", true).
		Order("solves.is_correct desc").
		Limit(1).Find(&task).Error

	if err != nil {
		log.Println(err)
		return TaskWithGenre{}
	}

	return task
}

func DeleteTask(dbsess *gorm.DB, taskID int) bool {

	err := dbsess.Table("tasks").Where("tasks.id = ?", taskID).Delete(Task{}).Error

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func UpdateTask(dbsess *gorm.DB, task *Task) bool {

	err := dbsess.Table("tasks").Save(&task).Error

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
