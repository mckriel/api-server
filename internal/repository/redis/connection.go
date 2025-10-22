package redis

import (
	"api-servers/internal"
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

type Database struct {
	Connection *redis.Client
}

var (
	db_instance *Database
	once        sync.Once
)

func GetDatabase() (*Database, error) {
	var err error

	once.Do(func() {
		connection_string := internal.CONN_REDIS
		client := redis.NewClient(&redis.Options{
			Addr: connection_string,
		})

		ctx := context.Background()
		if ping_error := client.Ping(ctx).Err(); ping_error != nil {
			err = ping_error
			return
		}

		db_instance = &Database{
			Connection: client,
		}
	})

	return db_instance, err
}

func CloseDatabase() error {
	if db_instance != nil && db_instance.Connection != nil {
		return db_instance.Connection.Close()
	}
	return nil
}
