package main

import (
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

type Database interface {
	Initialize() error
	Disconnect() error
}

type DataResult interface {
	UserPermissionResult
}

type UserPermissionResult struct {
	Name           string
	PermissionName string
	ObjectName     *string
}

type QueryResult[T DataResult] struct {
	Duration time.Duration
	Data     []T
}
