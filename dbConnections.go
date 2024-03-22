package main

import (
	"context"
	"database/sql"
	"log"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)

type QueryResult interface {
	UserPermissionResult
}

type UserPermissionResult struct {
	name           string
	permissionName string
	objectName     *string
}

type DbConnection[K QueryResult] interface {
	connect() error
	disconnect() error
	query(ctx context.Context, query string, parameters *map[string]interface{}) ([]K, error)
	nonQuery(ctx context.Context, query string, parameters *map[string]interface{}) error
}

type MsSqlConnection[K QueryResult] struct {
	server   string
	database string
	username string
	password string

	connection *sql.DB
}

func (m *MsSqlConnection[K]) connect() error {
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

func (m *MsSqlConnection[K]) disconnect() error {
	return m.connection.Close()
}

func (m *MsSqlConnection[K]) query(ctx context.Context, query string, parameters *map[string]interface{}) ([]K, error) {
	// Generate the parameter list.
	queryParams := make([]sql.NamedArg, len(*parameters))
	index := 0
	for key, value := range *parameters {
		queryParams[index] = sql.Named(key, value)
		index++
	}
	// Make the query
	result, err := m.connection.QueryContext(ctx, query, queryParams)
	if err != nil {
		return nil, err
	}
	// Create an output and iterate over the result adding to the output.
	output := make([]K, 0)
	index = 0
	for result.Next() {
		x := K{}
		err = result.Scan(&x)
		if err != nil {
			log.Printf("Error in scan:%e\n", err)
			continue
		}
		output = append(output, x)
	}
	return output, nil
}

func (m *MsSqlConnection[K]) nonQuery(ctx context.Context, query string, parameters *map[string]interface{}) error {
	queryParams := make([]sql.NamedArg, len(*parameters))
	index := 0
	for key, value := range *parameters {
		queryParams[index] = sql.Named(key, value)
		index++
	}
	_, err := m.connection.ExecContext(ctx, query, queryParams)
	if err != nil {
		return err
	}
	return nil
}
