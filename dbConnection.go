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
<<<<<<< HEAD
	FindUsers(target string) (QueryResult[UserPermissionResult], error)
	GrantPermissions(user string, target string, permission string) (bool, error)
	RemovePermission(user string, target string, permission string) (bool, error)
}

type UserPermissionResult struct {
	ID             *string `json:"_id"`
	Name           string  `json:"Name"`
	PermissionName string  `json:"PermissionName"`
	ObjectName     *string `json:"ObjectName"`
=======
}

type UserPermissionResult struct {
	Name           string  `json:"name"`
	PermissionName string  `json:"permission_name"`
	ObjectName     *string `json:"object_name"`
>>>>>>> c4c1b26fd7d0469b658074c0b37f9421346e5d22
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
