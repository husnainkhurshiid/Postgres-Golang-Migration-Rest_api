package main

import (
	"products/database"
	"products/migrator"
)

func main() {
	db := database.Connection()
	migrator.Migrate(db)
	defer db.Close()
}
