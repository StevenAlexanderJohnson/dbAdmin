package main

import (
	"context"
	"database/sql"
	"log"
	"net/url"
	"time"
)

type MsSqlDatabase struct {
	Database
	server   string
	database string
	username string
	password string

	ctx        context.Context
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

func (m *MsSqlDatabase) QueryUserPermissions(user string, target string) (QueryResult[UserPermissionResult], error) {
	tsql := `
	SELECT p.name, dp.permission_name, o.name
	FROM sys.database_principles p
	JOIN sys.database_permissions dp on dp.grantee_principal_id = p.principal_id
	LEFT JOIN sys.objects o on o.object_id = dp.major_id
	WHERE p.name = @user and (@target = '' or o.name = @target)
	`
	output := QueryResult[UserPermissionResult]{
		duration: time.Since(time.Now()),
		data:     nil,
	}
	outputData := make([]UserPermissionResult, 0)
	rows, err := m.connection.QueryContext(m.ctx, tsql, sql.Named("@user", user), sql.Named("@target", target))
	if err != nil {
		output.data = nil
		return output, err
	}
	for rows.Next() {
		temp := UserPermissionResult{}
		err = rows.Scan(&temp)
		if err != nil {
			log.Println("Error reading row from User Permissions result.")
		}
		outputData = append(outputData, temp)
	}
	output.data = outputData
	return output, nil
}
