package main

import (
	"database/sql"
	"gee-orm/log"
	"gee-orm/session"
)

type Engine struct {
	db *sql.DB
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
	e := &Engine{db: db}
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
	return session.NewSession(e.db)
}
