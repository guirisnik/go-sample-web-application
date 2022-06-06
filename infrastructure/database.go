package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

var dbInstance *sql.DB
var dbConnectionError error

func GetInstance() *sql.DB {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			fmt.Println("Connecting to postgres database local_db.")
			dbInstance, dbConnectionError = sql.Open("postgres", "postgres://local_user:pg_local_password@localhost:5432/local_db?sslmode=disable")
			if dbConnectionError != nil {
				log.Fatal(dbConnectionError)
			} else {
				fmt.Println("Connection successful.")
			}
		} else {
			fmt.Println("Database instance already exists.")
		}
	} else {
		fmt.Println("Database instance already exists.")
	}

	return dbInstance
}
