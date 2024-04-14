package main

import (
	"context"
	"database/sql"
	"encoding/json"
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
	localDb      SqlLiteDatabase
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
	// Initialize SQLite db if it isn't already initialized.
	appDataDir := filepath.Join(os.Getenv("APPDATA"), "dbAdmin")
	_, err := os.Stat(appDataDir)
	if os.IsNotExist(err) {
		// 0755 means that owner an read/write/execute but all other applications can only read.
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
	localDb, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatalf("Unable to connect to app data db.\n%e\n", err)
	}
	a.localDb = SqlLiteDatabase{
		connection: localDb,
		ctx:        a.ctx,
	}
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	log.Println("######### THE DOM IS READY #########")
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
	log.Println(server, database, driver, username, password)
	databaseKey := fmt.Sprintf("%s:%s", server, database)
	if _, ok := a.databaseHash[databaseKey]; ok {
		return fmt.Sprintf("%s has already been registered.", databaseKey)
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
			sqlite:   &a.localDb,
		}
	case "mongo":
		connection = &MongoDatabase{
			server:   server,
			username: username,
			password: password,
			ctx:      a.ctx,
			sqlite:   &a.localDb,
		}
	default:
		return fmt.Sprintf("Invalid driver was selected: %s\n", driver)
	}
	if err := connection.Initialize(); err != nil {
		return fmt.Sprintf("There was an error connecting to the database.\n%e\n", err)
	}
	a.databaseHash[databaseKey] = connection
	return "Successfully connected to the database."
}

func (a *App) GetUserPermissions(databaseKey string, user string, target string) (string, error) {
	db, ok := a.databaseHash[databaseKey]
	if !ok {
		return "", fmt.Errorf("%s has not been registered yet", databaseKey)
	}

	queryResult, err := db.FindUserPermissions(user, target)
	if err != nil {
		return "", fmt.Errorf("an error occurred while collecting user permissions\n%s", err)
	}
	output, err := json.Marshal(queryResult)
	return string(output), err
}

func (a *App) GetConnections() []string {
	output := make([]string, 0, len(a.databaseHash))

	for k := range a.databaseHash {
		output = append(output, k)
	}
	return output
}
