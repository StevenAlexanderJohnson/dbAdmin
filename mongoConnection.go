package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase interface {
	Database
	FindUserPermissions() (QueryResult[UserPermissionResult], error)
}

type MongoConnection struct {
	server   string
	username string
	password string

	connection *mongo.Client
}

func (m *MongoConnection) Initialize() error {
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
		return err
	}
	m.connection = client
	return nil
}

func (m *MongoConnection) Disconnect() error {
	return nil
}

func (m *MongoConnection) GetUserPermissions() (QueryResult[UserPermissionResult], error) {
	return QueryResult[UserPermissionResult]{
		duration: time.Since(time.Now()),
		data:     nil,
	}, nil
}
