package main

import (
	"database/sql"
	"net/url"
	"time"
)

type MsSqlDatabase struct {
	Database
	server   string
	database string
	username string
	password string

	connection *sql.DB
}

func (m *MsSqlDatabase) Initialize() error {
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

func (m *MsSqlDatabase) Disconnect() error {
	if err := m.connection.Close(); err != nil {
		return err
	}
	return nil
}

func (m *MsSqlDatabase) Connection() *sql.DB {
	return m.connection
}

func QueryUserPermissions(database *MsSqlDatabase) (QueryResult[UserPermissionResult], error) {
	return QueryResult[UserPermissionResult]{
		duration: time.Since(time.Now()),
		data:     nil,
	}, nil
}
