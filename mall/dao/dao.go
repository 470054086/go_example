package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var G_db *gorm.DB

type Dao struct {
}

func NewDao(host string) {
	fmt.Println(host)
	db, err := gorm.Open("mysql", host)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	G_db = db
}
