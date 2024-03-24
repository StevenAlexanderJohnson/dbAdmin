package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func (m *MongoConnection) Connection() *mongo.Client {
	return m.connection
}
