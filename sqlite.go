package main

import (
	"context"
	"database/sql"
	"time"
)

type Severity int

const (
	LOG     Severity = 0
	WARNING Severity = 1
	ERROR   Severity = 2
)

type SqlLiteDatabase struct {
	connection *sql.DB
	ctx        context.Context
}

func (s *SqlLiteDatabase) Initialize() error {
	_, err := s.connection.ExecContext(
		s.ctx,
		`CREATE TABLE IF NOT EXISTS Logs(
			id INTEGER NOT NULL PRIMARY KEY,
			severity INTEGER NOT NULL,
			error_message TEXT NOT NULL,
			error_file TEXT NOT NULL,
			error_function TEXT NOT NULL,
			date_created TEXT NOT NULL
		)`,
	)
	return err
}

func (s *SqlLiteDatabase) WriteLog(severity Severity, logError error, errorFile string, errorFunction string) error {
	_, err := s.connection.ExecContext(
		s.ctx,
		`INSERT INTO Logs (severity, error_message, error_file, error_function, date_created) VALUES (?, ?, ?, ?, ?)`,
		severity, logError, errorFile, errorFunction, time.Now().UTC(),
	)
	return err
}
