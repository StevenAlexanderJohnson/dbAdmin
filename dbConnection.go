package main

import (
	"database/sql"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

type Database interface {
	Initialize() error
	Disconnect() error
	Connection() *sql.DB
}

type DataResult interface {
	UserPermissionResult
}

type UserPermissionResult struct {
	name           string
	permissionName string
	objectName     *string
}

type QueryResult[T DataResult] struct {
	duration time.Duration
	data     []T
}
