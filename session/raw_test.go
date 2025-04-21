package session

import (
	"database/sql"
	"gee-orm/log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSession_Exec(t *testing.T) {
	TestDB, err := sql.Open("sqlite3", "/Users/bytedance/gee-orm/gee.db")
	if TestDB == nil || err != nil {
		log.Error(err)
		return
	}
	defer TestDB.Close()
	s := NewSession(TestDB)
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2, but got", count)
	}
}

func TestSession_QueryRows(t *testing.T) {
	TestDB, err := sql.Open("sqlite3", "/Users/bytedance/gee-orm/gee.db")
	if TestDB == nil || err != nil {
		log.Error(err)
	}
	defer TestDB.Close()
	s := NewSession(TestDB)
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	row := s.Raw("SELECT count(*) FROM User").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatal("failed to query db", err)
	}
}
