package main

import (
	"context"
	"database/sql"
	"log"
	"net/url"
	"time"
)

type MsSqlDatabase struct {
	server   string
	database string
	username string
	password string

	ctx        context.Context
	connection *sql.DB
	sqlite     *SqlLiteDatabase
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
		m.sqlite.WriteLog(ERROR, err, "msSqlConnection.go", "Initialize")
		return err
	}
	m.connection = db
	return db.PingContext(m.ctx)
}

func (m *MsSqlDatabase) Disconnect() error {
	if err := m.connection.Close(); err != nil {
		m.sqlite.WriteLog(ERROR, err, "msSqlConnection.go", "Disconnect")
		return err
	}
	return nil
}

func (m *MsSqlDatabase) FindUsers(target string) (QueryResult[UserPermissionResult], error) {
	tsql := `
	SELECT distinct p.name
	FROM sys.database_principals p
	JOIN sys.database_permissions dp on dp.grantee_principal_id = p.principal_id
	LEFT JOIN sys.objects o on o.object_id = dp.major_id
	`
	output := QueryResult[UserPermissionResult]{
		Duration: time.Since(time.Now()),
		Data:     nil,
	}
	outputData := make([]UserPermissionResult, 0)
	rows, err := m.connection.QueryContext(m.ctx, tsql)
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "msSqlConnection.go", "QueryUserPermissions")
		output.Data = nil
		return output, err
	}
	for rows.Next() {
		temp := UserPermissionResult{}
		err = rows.Scan(&temp.Name)
		if err != nil {
			m.sqlite.WriteLog(ERROR, err, "msSqlConnection.go", "QueryUserPermissions")
			log.Println("Error reading row from User Permissions result.")
		}
		outputData = append(outputData, temp)
	}
	output.Data = outputData
	return output, nil
}

func (m *MsSqlDatabase) FindUserPermissions(user string) (QueryResult[UserPermissionResult], error) {
	tsql := `
	SELECT p.name, dp.permission_name, o.name
	FROM sys.database_principals p
	JOIN sys.database_permissions dp on dp.grantee_principal_id = p.principal_id
	LEFT JOIN sys.objects o on o.object_id = dp.major_id
	WHERE p.name = @user
	`
	output := QueryResult[UserPermissionResult]{
		Duration: time.Since(time.Now()),
		Data:     nil,
	}
	outputData := make([]UserPermissionResult, 0)
	rows, err := m.connection.QueryContext(m.ctx, tsql, sql.Named("user", user))
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "msSqlConnection.go", "QueryUserPermissions")
		output.Data = nil
		return output, err
	}
	for rows.Next() {
		temp := UserPermissionResult{}
		err = rows.Scan(&temp.Name, &temp.PermissionName, &temp.ObjectName)
		if err != nil {
			m.sqlite.WriteLog(ERROR, err, "msSqlConnection.go", "QueryUserPermissions")
		}
		outputData = append(outputData, temp)
	}
	output.Data = outputData
	return output, nil
}

func (m *MsSqlDatabase) GrantPermissions(user string, target string, permission string) (bool, error) {
	tsql := `GRANT @perm ON @target TO @user`
	res, err := m.connection.ExecContext(m.ctx, tsql, sql.Named("perm", permission), sql.Named("user", user), sql.Named("target", target))
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "mssqlConnection.go", "GrantPermissions")
		return false, err
	}
	log.Println(res)
	return true, nil
}

func (m *MsSqlDatabase) RemovePermission(user string, target string, permission string) (bool, error) {
	tsql := `REVOKE @perm ON @target TO @user`
	log.Println(user, target, permission)
	res, err := m.connection.ExecContext(m.ctx, tsql, sql.Named("perm", permission), sql.Named("user", user), sql.Named("target", target))
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "mssqlConnection.go", "GrantPermissions")
		return false, err
	}
	log.Println(res)
	return true, nil
}
