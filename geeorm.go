package main

import (
	"database/sql"
	"fmt"
	"gee-orm/dialect"
	"gee-orm/log"
	"gee-orm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (*Engine, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err = db.Ping(); err != nil { // Send a ping to make sure the database connection is alive.
		log.Error(err)
		return nil, err
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s not exist", driver)
		return nil, fmt.Errorf("dialect %s not support", driver)
	}
	e := &Engine{
		db:      db,
		dialect: dial,
	}
	log.Info("Connected to database successfully")
	return e, nil
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
		return
	}
	log.Info("Closing database successfully")
}

func (e *Engine) NewSession() *session.Session {
	return session.NewSession(e.db, e.dialect)
}
