package migrator

import (
	"database/sql"
	"flag"
	"fmt"
	"products/shared"

	"github.com/gobuffalo/packr"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func Migrate(db *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("../migration"),
	}

	flagValue := flag.String("migration", "", "Migration File")
	flag.Parse()

	switch *flagValue {
	case "up":
		migrate, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		shared.CheckError(err)
		fmt.Printf("\n%d Migration applied", migrate)
	case "down":
		migrate, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
		shared.CheckError(err)
		fmt.Printf("\n%d Migration applied", migrate)
	}
}
