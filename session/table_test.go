package session

import (
	"gee-orm/log"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	if TestDB == nil || err != nil {
		log.Error(err)
		return
	}
	defer TestDB.Close()
	s := NewSession(TestDB, dia).Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
