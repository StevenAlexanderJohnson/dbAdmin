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
	ctx      context.Context

	connection *mongo.Client
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
		return err
	}
	m.connection = client
	return client.Ping(m.ctx, nil)
}

func (m *MongoDatabase) Disconnect() error {
	return m.connection.Disconnect(context.TODO())
}

func (m *MongoDatabase) FindUserPermissions() (QueryResult[UserPermissionResult], error) {
	return QueryResult[UserPermissionResult]{
		Duration: time.Since(time.Now()),
		Data:     nil,
	}, nil
}
