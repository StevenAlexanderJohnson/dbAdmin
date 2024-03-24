package main

import (
	"context"
	"database/sql"
	"log"
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

func Query[T DataResult](db Database, ctx context.Context, query string, parameters *map[string]interface{}) (QueryResult[T], error) {
	queryParams := make([]sql.NamedArg, len(*parameters))
	index := 0
	for key, value := range *parameters {
		queryParams[index] = sql.Named(key, value)
		index++
	}

	start := time.Now()
	result, err := db.Connection().QueryContext(ctx, query, queryParams)
	output := QueryResult[T]{
		duration: time.Since(start),
	}
	if err != nil {
		return output, err
	}
	dataOutput := make([]T, 0)
	for result.Next() {
		x := T{}
		err = result.Scan(&x)
		if err != nil {
			log.Printf("Error in scanning query result:%e\n", err)
			continue
		}
		dataOutput = append(dataOutput, x)
	}
	output.data = dataOutput
	return output, nil
}
