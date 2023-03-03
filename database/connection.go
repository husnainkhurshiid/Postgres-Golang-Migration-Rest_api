package database

import (
	"database/sql"
	"fmt"
)

func Connection() (*sql.DB, error) {
	connectionString := "postgres://postgres:postgres@localhost/shop?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	CheckError(err)
	return db, nil
}

func CheckError(err error) {
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Database Connected")
	}
}
