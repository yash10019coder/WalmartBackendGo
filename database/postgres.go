package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_NAME     = "postgres"
	DB_USERNAME = "postgres"
	DB_PASSWORD = "postgres"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
)

func SetupDB() *sql.DB {
	var dbinfo string = fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	return db
}
