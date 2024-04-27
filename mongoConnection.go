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

<<<<<<< HEAD
func (m *MongoDatabase) FindUsers(target string) (QueryResult[UserPermissionResult], error) {
	var output QueryResult[UserPermissionResult] = QueryResult[UserPermissionResult]{}
	startTime := time.Now()
	db := m.connection.Database(target)
	if db == nil {
		m.sqlite.WriteLog(ERROR, fmt.Errorf("user tried accessing a database that didn't exists"), "mongoConnection.go", "FindUsers:Database()")
		return output, nil
	}
	col := db.Collection("system.users")
	if col == nil {
		m.sqlite.WriteLog(ERROR, fmt.Errorf("system.users table is not available for the database selected"), "mongoConnection.go", "FindUsers:Database()")
		return output, nil
	}
	cursor, err := col.Distinct(m.ctx, "user", bson.D{})
	output.Duration = time.Since(startTime)
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "mongoConnection.go", "FindUsers")
		return output, err
	}

	var data []UserPermissionResult
	for _, result := range cursor {
		if s, ok := result.(string); ok {
			data = append(data, UserPermissionResult{
				Name:           s,
				PermissionName: "",
				ObjectName:     nil,
			})
		}
	}
	output.Data = data
	return output, nil
}

func (m *MongoDatabase) FindUserPermissions(user string, target string) (QueryResult[UserPermissionResult], error) {
	fmt.Printf("Searching for %s within %s", user, target)
	var output QueryResult[UserPermissionResult]
	startTime := time.Now()
	cursor, err := m.connection.Database("admin").Collection("system.users").Aggregate(m.ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "user", Value: user}, {Key: "roles.db", Value: bson.D{{Key: "$regex", Value: target}}}}}},
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$roles"}}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "Name", Value: "$user"},
			{Key: "PermissionName", Value: "$roles.role"},
			{Key: "ObjectName", Value: "$roles.db"},
=======
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
>>>>>>> c4c1b26fd7d0469b658074c0b37f9421346e5d22
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
<<<<<<< HEAD
			continue
=======
>>>>>>> c4c1b26fd7d0469b658074c0b37f9421346e5d22
		}
		data = append(data, result)
	}
	if err := cursor.Err(); err != nil {
		m.sqlite.WriteLog(ERROR, err, "mongoConnection.go", "FindUserPermissions:cursor.Err()")
<<<<<<< HEAD
		return output, err
	}
	output.Data = data
	return output, nil
}

func (m *MongoDatabase) GrantPermissions(user string, target string, permission string) (bool, error) {
	cursor, err := m.connection.Database("admin").
		Collection("system.users").
		UpdateOne(
			m.ctx,
			bson.D{{Key: "_id", Value: user}},
			bson.D{{Key: "$push", Value: bson.D{{Key: "roles", Value: bson.D{{Key: "role", Value: permission}, {Key: "db", Value: target}}}}}},
		)
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "mongoConnection", "GrantPermissions:UpdateOne")
		return false, err
	}
	if cursor.ModifiedCount == 0 {
		return false, fmt.Errorf("user was not able to be found using the provided data")
	}
	if cursor.ModifiedCount > 1 || cursor.MatchedCount > 1 {
		return true, fmt.Errorf("more than one user was found using %s. number of updates: %d. matched records: %d", user, cursor.ModifiedCount, cursor.MatchedCount)
	}
	return true, nil
}

func (m *MongoDatabase) RemovePermission(user string, target string, permission string) (bool, error) {
	cursor, err := m.connection.Database("admin").Collection("system.user").UpdateOne(
		m.ctx,
		bson.D{{Key: "_id", Value: user}},
		bson.D{{Key: "$pull", Value: bson.D{{Key: "roles", Value: bson.D{{Key: "role", Value: permission}, {Key: "db", Value: target}}}}}},
	)
	if err != nil {
		m.sqlite.WriteLog(ERROR, err, "mongoConnection", "RemovePermission:UpdateOne")
		return false, err
	}
	if cursor.ModifiedCount == 0 {
		return false, fmt.Errorf("user was not able to be found using the provided data")
	}
	if cursor.ModifiedCount > 1 || cursor.MatchedCount > 1 {
		return true, fmt.Errorf("more than one user was found using %s. number of updates: %d. matched records: %d", user, cursor.ModifiedCount, cursor.MatchedCount)
	}
	return true, nil
=======
		return output, nil
	}
	output.Data = data
	return output, nil
>>>>>>> c4c1b26fd7d0469b658074c0b37f9421346e5d22
}
