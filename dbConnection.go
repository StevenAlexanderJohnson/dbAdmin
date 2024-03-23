package main

import (
	"context"
	"database/sql"
	"log"
	"net/url"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

type Database interface {
	Initialize() error
	Disconnect() error
	Connection() *sql.DB
}

type DataResult interface {
	UserPermissionResult2
}

type UserPermissionResult2 struct {
	name           string
	permissionName string
	objectName     *string
}

type QueryResult2[T DataResult] struct {
	duration time.Duration
	data     []T
}

type MsSqlConnection2 struct {
	server   string
	database string
	username string
	password string

	connection *sql.DB
}

func (m *MsSqlConnection2) Initialize() error {
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

func (m *MsSqlConnection2) Disconnect() error {
	if err := m.connection.Close(); err != nil {
		return err
	}
	return nil
}

func (m *MsSqlConnection2) Connection() *sql.DB {
	return m.connection
}

func Query[T DataResult](db Database, ctx context.Context, query string, parameters *map[string]interface{}) (QueryResult2[T], error) {
	queryParams := make([]sql.NamedArg, len(*parameters))
	index := 0
	for key, value := range *parameters {
		queryParams[index] = sql.Named(key, value)
		index++
	}

	start := time.Now()
	result, err := db.Connection().QueryContext(ctx, query, queryParams)
	output := QueryResult2[T]{
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
