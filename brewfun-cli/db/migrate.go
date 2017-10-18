package db

import (
	"database/sql"
	"fmt"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/sqlite3"
	_ "github.com/mattes/migrate/source/file"
	"github.com/urfave/cli"
)

const (
	databaseName = "sqlite3"
)

func Migrate() *cli.ExitError {
	db, err := sql.Open(databaseName, "test.sqlite3?_foreign_keys=1")
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		err = fmt.Errorf("%s test.sqlite3", err)
		return cli.NewExitError(err, 1)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations",
		databaseName, driver)
	if err != nil {
		return cli.NewExitError(err, 2)
	}

	err = m.Up()
	if err == nil {
		return nil
	}
	if err == migrate.ErrNoChange {
		return nil
	}
	return cli.NewExitError(err, 3)
}
