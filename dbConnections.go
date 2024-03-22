package main

import (
	"database/sql"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)

type DbConnection interface {
	connect() error
	disconnect() error
}

type MsSqlConnection struct {
	server   string
	database string
	username string
	password string

	connection *sql.DB
}

func (m *MsSqlConnection) connect() error {
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

func (m *MsSqlConnection) disconnect() error {
	return m.connection.Close()
}
