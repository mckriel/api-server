package internal

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONN_MYSQL   = "root:password@tcp(localhost:3306)/api_mysql?parseTime=true"
	CONN_MONGODB = "mongodb://localhost:27017/api_mongodb"
	CONN_REDIS   = "localhost:6379"
)

func GetMySQLConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", CONN_MYSQL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GetMongoDBConnection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONN_MONGODB))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
