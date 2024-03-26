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
	Name           string  `json:"Name"`
	PermissionName string  `json:"PermissionName"`
	ObjectName     *string `json:"ObjectName"`
}

type QueryResult[T DataResult] struct {
	Duration time.Duration `json:"Duration"`
	Data     []T           `json:"Data"`
}
