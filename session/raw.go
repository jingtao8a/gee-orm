package session

import (
	"database/sql"
	"gee-orm/dialect"
	"gee-orm/log"
	"gee-orm/schema"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	sql      strings.Builder
	sqlVars  []interface{}
}

func NewSession(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 执行INSERT/UPDATE/DELETE
func (s *Session) Exec() (sql.Result, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	result, err := s.db.Exec(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
	}
	return result, err
}

// 执行SELECT
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.db.QueryRow(s.sql.String(), s.sqlVars...) // 返回单个Row对象，延迟到调用Scan()时返回错误，自动关闭底层连接
}

func (s *Session) QueryRows() (*sql.Rows, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	rows, err := s.db.Query(s.sql.String(), s.sqlVars...) // 返回Rows迭代器，需要手动调用rows.close()
	if err != nil {
		log.Error(err)
	}
	return rows, err
}
