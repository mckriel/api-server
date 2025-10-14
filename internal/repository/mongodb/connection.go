package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var db_instance *Database

func Connect() (*Database, error) {
	if db_instance != nil {
		return db_instance, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connection_uri := "mongodb://root:password@localhost:27017"
	client_options := options.Client().ApplyURI(connection_uri)

	client, err := mongo.Connect(ctx, client_options)
	if err != nil {
		return nil, fmt.Errorf("mongodb connection failed: %w", err)
	}

	database := client.Database("api_mongodb")

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	db_instance = &Database{
		Client:   client,
		Database: database,
	}

	log.Println("Mongodb connection successful")
	return db_instance, nil
}

func (d *Database) Disconnect() error {
	if d.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return d.Client.Disconnect(ctx)
	}
	return nil
}
