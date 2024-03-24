package main

import (
	"context"
	"database/sql"
	"log"
	"net/url"
	"time"
)

type MsSqlConnection struct {
	server   string
	database string
	username string
	password string

	connection *sql.DB
}

func (m *MsSqlConnection) Initialize() error {
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

func (m *MsSqlConnection) Disconnect() error {
	if err := m.connection.Close(); err != nil {
		return err
	}
	return nil
}

func (m *MsSqlConnection) Connection() *sql.DB {
	return m.connection
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
