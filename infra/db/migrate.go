package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {

	migratements := &migrate.FileMigrationSource{
		Dir: dir,
	}

	_, err := migrate.Exec(db.DB, "postgres", migratements, migrate.Up)
	if err != nil {
		return err
	}

	fmt.Println("Database migration completed successfully")
	return nil
}