package mysql

import (
	"api-servers/internal"
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connection *sql.DB
}

var (
	db_instance *Database
	once        *sync.Once
)

func GetDatabase() (*Database, error) {
	var err error

	once.Do(func() {
		connection_string := internal.CONN_MONGODB
		conn, connection_error := sql.Open("mysql", connection_string)

		if connection_error != nil {
			err = connection_error
			return
		}

		if ping_error := conn.Ping(); ping_error != nil {
			err = ping_error
			return
		}

		db_instance = &Database{
			Connection: conn,
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
