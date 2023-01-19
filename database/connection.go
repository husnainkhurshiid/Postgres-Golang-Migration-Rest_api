package database

import (
	"database/sql"
	"fmt"
)

func Connection() *sql.DB {
	connectionString := "postgres://postgres:postgres@localhost/shop?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	CheckError(err)
	return db
}

func CheckError(err error) {
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Database Connection Eastablished")
	}
}
