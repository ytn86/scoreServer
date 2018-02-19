package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	dbsession *gorm.DB
)
	
type DBController struct {
	*revel.Controller
	Tx *gorm.DB // Transaction
}

func (c *DBController) Begin() revel.Result {
	c.Tx = dbsession.Begin()

	return nil
}

func (c *DBController) Commit() revel.Result {
	if c.Tx == nil {
		return nil
	}
	c.Tx.Commit()
	c.Tx = nil

	return nil
}

func (c *DBController) RollBack() revel.Result {
	if c.Tx == nil{
		return nil
	}
	c.Tx.Rollback()
	c.Tx = nil

	return nil
}
	

func InitDB() {

	var err error
	dbPath := "./scoreServer.db"
	dbsession, err = gorm.Open("sqlite3", dbPath)

	if err != nil {
		panic(err)
	}

}

func CloseDB() {
	dbsession.Close()
}
