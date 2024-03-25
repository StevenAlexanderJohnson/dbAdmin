package main

import (
	"database/sql"
	"net/url"
)

type MsSqlDatabase interface {
	Database
	Query() error
}

type MsSqlConnection struct {
	server   string
	database string
	username string
	password string

	connection *sql.DB
}

func (m *MsSqlConnection) Initialize() error {
	query := url.Values{}
	query.Add("database", m.database)
	url := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(m.username, m.password),
		Host:     m.server,
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", url.String())
	if err != nil {
		return err
	}
	m.connection = db
	return nil
}

func (m *MsSqlConnection) Disconnect() error {
	if err := m.connection.Close(); err != nil {
		return err
	}
	return nil
}

func (m *MsSqlConnection) Connection() *sql.DB {
	return m.connection
}

func (m *MsSqlConnection) Query() error {
	return nil
}
