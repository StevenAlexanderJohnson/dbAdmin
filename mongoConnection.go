package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	server   string
	username string
	password string

	ctx        context.Context
	connection *mongo.Client
	sqlite     *SqlLiteDatabase
}

func (m *MongoDatabase) Initialize() error {
	fmt.Println("MONGO", m.server, m.username, m.password)
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(
			fmt.Sprintf(
				"mongodb://%s:%s@%s/?authMechanism=DEFAULT",
				m.username,
				m.password,
				m.server,
			),
		),
	)
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "mongoConnection.go", "Initialize")
		return err
	}
	m.connection = client
	return client.Ping(m.ctx, nil)
}

func (m *MongoDatabase) Disconnect() error {
	err := m.connection.Disconnect(context.TODO())
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "mongoConnection.go", "Disconnect")
		return err
	}
	return nil
}

func (m *MongoDatabase) FindUserPermissions() (QueryResult[UserPermissionResult], error) {
	var output QueryResult[UserPermissionResult]
	startTime := time.Now()
	output = QueryResult[UserPermissionResult]{
		Duration: time.Since(startTime),
		Data:     nil,
	}
	// if err != nil {
	// 	m.sqlite.WriteLog(ERROR, err, "mongoConnection.go", "FindUserPermissions")
	// 	return output, err
	// }
	return output, nil
}
