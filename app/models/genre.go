package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Genre struct {
	ID    int    `gorm:"primary_key" json:"id,omitempty"`
	Name  string `sql:"unique" json:"name,omitempty"`
	Color string `sql:"unique" json:"color,omitempty"`
}

func GetAllGenresOrderByIDAsc(dbsess *gorm.DB) []Genre {

	var genres []Genre
	dbsess.Table("genres").Order("id asc").Find(&genres)

	return genres
}

func AddGenre(dbsess *gorm.DB, genre *Genre) bool {

	err := dbsess.Create(&genre).Error
	if err != nil {
		return false
	}
	
	return true
}

func UpdateGenre(dbsess *gorm.DB, genre *Genre) bool {

	err := dbsess.Table("genres").Save(genre).Error

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func DeleteGenreFilterByID(dbsess *gorm.DB, genreID int) bool {

	if dbsess.Table("genres").Where("id = ?", genreID).Delete(Genre{}).Error != nil {
		return false
	}
	return true
}
