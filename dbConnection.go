package main

import (
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

// Interface all Database structs (i.e. MongoDatabase) should fulfil.
type Database interface {
	Initialize() error
	Disconnect() error
	FindUserPermissions(user string, target string) (QueryResult[UserPermissionResult], error)
	FindUsers(target string) (QueryResult[UserPermissionResult], error)
	GrantPermissions(user string, target string, permission string) (bool, error)
}

type UserPermissionResult struct {
	Name           string  `json:"Name"`
	PermissionName string  `json:"PermissionName"`
	ObjectName     *string `json:"ObjectName"`
}

// This interface acts like a constraint for what structs can be used for the generic type T in QueryResult.
// Any struct that is used to parse what's returned from queries should be added here.
type DataResult interface {
	UserPermissionResult
}

// Struct that will be returned from Go to the frontend.
type QueryResult[T DataResult] struct {
	Duration time.Duration `json:"Duration"`
	Data     []T           `json:"Data"`
}
