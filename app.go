package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// App struct
type App struct {
	ctx          context.Context
	databaseHash map[string]Database
	localDb      *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	a.databaseHash = make(map[string]Database)

	// Connect to the SQLite database.
	appDataDir := filepath.Join(os.Getenv("APPDATA"), "dbAdmin")
	_, err := os.Stat(appDataDir)
	// Create the app data folder if it doesn't exists.
	if os.IsNotExist(err) {
		err := os.MkdirAll(appDataDir, 0755)
		if err != nil {
			log.Fatalf("Unable to create app data directory.\n%e\n", err)
		}
	}
	// Create the db file if it doesn't exists.
	databasePath := filepath.Join(appDataDir, "data.db")
	_, err = os.Stat(databasePath)
	if os.IsNotExist(err) {
		_, err = os.Create(databasePath)
		if err != nil {
			log.Fatalf("Unable to find or create app data db.\n%e\n", err)
		}
	}
	a.localDb, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatalf("Unable to connect to app data db.\n%e\n", err)
	}
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	for _, connection := range a.databaseHash {
		err := connection.Disconnect()
		if err != nil {
			log.Println(err)
		}
	}
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Registers a database connection in memory.
//
// Later on this will also register into a SQLite local DB to save credentials and connection information.
func (a *App) RegisterDatabase(server string, database string, driver string, username string, password string) string {
	switch driver {
	case "mssql":
		databaseKey := fmt.Sprintf("%s:%s", server, database)
		if _, ok := a.databaseHash[databaseKey]; ok {
			return fmt.Sprintf("%s has already been registered.\n", databaseKey)
		}
		var connection Database
		switch driver {
		case "mssql":
			connection = &MsSqlDatabase{
				server:   server,
				database: database,
				username: username,
				password: password,
				ctx:      a.ctx,
			}
		case "mongo":
			connection = &MongoDatabase{
				server:   server,
				username: username,
				password: password,
			}
		}
		if err := connection.Initialize(); err != nil {
			return fmt.Sprintf("There was an error connecting to the database.\n%e\n", err)
		}
		a.databaseHash[databaseKey] = connection

	default:
		return fmt.Sprintf("%s is not a valid driver.\n", driver)
	}
	return "Successfully connected to the database."
}

func (a *App) GetUserPermissions(databaseKey string, user string, target string) (QueryResult[UserPermissionResult], error) {
	db, ok := a.databaseHash[databaseKey]
	if !ok {
		return QueryResult[UserPermissionResult]{}, fmt.Errorf("%s has not been registered yet", databaseKey)
	}

	var err error
	var queryResult QueryResult[UserPermissionResult]

	switch v := db.(type) {
	case *MongoDatabase:
		queryResult, err = v.FindUserPermissions()
	case *MsSqlDatabase:
		queryResult, err = v.QueryUserPermissions(user, target)
	default:
		return QueryResult[UserPermissionResult]{}, fmt.Errorf("an error occurred while collecting user permissions")
	}

	if err != nil {
		return QueryResult[UserPermissionResult]{}, fmt.Errorf("an error occurred while collecting user permissions\n%s", err)
	}
	return queryResult, nil
}
