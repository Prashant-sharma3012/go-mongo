package dbMongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conn *mongo.Database

type DB struct {
	IsMongo bool
	conn    *mongo.Database
}

func (db *DB) Connect() error {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return err
	}

	defer cancel()

	db.conn = client.Database("go-mdb")

	return nil
}

func (db *DB) GetConnection() *mongo.Database {
	return db.conn
}
