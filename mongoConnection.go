package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
				"mongodb://%s:%s@%s/?authMechanism=SCRAM-SHA-256",
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

func (m *MongoDatabase) FindUserPermissions(user string, target string) (QueryResult[UserPermissionResult], error) {
	var output QueryResult[UserPermissionResult]
	startTime := time.Now()
	cursor, err := m.connection.Database("admin").Aggregate(m.ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "user", Value: user}, {Key: "roles.db", Value: bson.D{{Key: "$regex", Value: target}}}}}},
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$roles"}}}},
		{{Key: "$project", Value: bson.D{
			{Key: "name", Value: 1},
			{Key: "role", Value: "$roles.role"},
			{Key: "db", Value: "$roles.db"},
		}}},
	})
	output = QueryResult[UserPermissionResult]{
		Duration: time.Since(startTime),
		Data:     nil,
	}
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "mongoConnection.go", "FindUserPermissions")
		return output, err
	}
	data := make([]UserPermissionResult, 0)
	for cursor.Next(m.ctx) {
		var result UserPermissionResult
		if err = cursor.Decode(&result); err != nil {
			m.sqlite.WriteLog(ERROR, err, "mongoConnection.go", "FindUserPermissions:cursor.Decode()")
		}
		data = append(data, result)
	}
	if err := cursor.Err(); err != nil {
		m.sqlite.WriteLog(ERROR, err, "mongoConnection.go", "FindUserPermissions:cursor.Err()")
		return output, nil
	}
	output.Data = data
	return output, nil
}
