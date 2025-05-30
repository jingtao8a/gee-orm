package main

import (
	"gee-orm/log"

	_ "github.com/mattn/go-sqlite3" // init sqlite3 database driver
)

type Hello struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func main() {
	engine, err := NewEngine("sqlite3", "/Users/bytedance/gee-orm/gee.db")
	if err != nil {
		log.Error(err)
	}
	session := engine.NewSession().Model(&Hello{})
	if session.HasTable() {
		session.DropTable()
	}
	session.CreateTable()
	res, err := session.Raw("INSERT INTO Hello(`Name`, `Age`) values (?, ?), (?,?)", "Tom", 12, "Sam", 13).Exec()
	if res == nil || err != nil {
		log.Error("INSERT FAILED1")
	}
	i, err := res.RowsAffected()
	if err != nil || i != 2 {
		log.Error("INSERT FAILED2")
	}
}
