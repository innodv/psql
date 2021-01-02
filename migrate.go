package psql

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/source/github"
	_ "github.com/lib/pq"
)

func MigrateDB(db *sql.DB, migrationDir string, dbName string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to get Postgres driver %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", migrationDir), dbName, driver)
	if err != nil {
		return fmt.Errorf("error running migrations %w", err)
	}
	err = m.Up()
	if err != nil {
		if err.Error() != "no change" {
			return fmt.Errorf("error applying migrations %w", err)
		}
	}
	return nil
}
