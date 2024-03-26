package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	Database
	server   string
	username string
	password string

	connection *mongo.Client
}

func (m *MongoDatabase) Initialize() error {
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

func (m *MongoDatabase) Disconnect() error {
	return m.connection.Disconnect(context.TODO())
}

func FindUserPermissions(database *MongoDatabase) (QueryResult[UserPermissionResult], error) {
	return QueryResult[UserPermissionResult]{
		duration: time.Since(time.Now()),
		data:     nil,
	}, nil
}
