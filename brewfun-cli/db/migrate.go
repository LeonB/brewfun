package db

import (
	"database/sql"
	"fmt"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/sqlite3"
	_ "github.com/mattes/migrate/source/file"
	"github.com/urfave/cli"
)

func Migrate() *cli.ExitError {
	db, err := sql.Open("sqlite3", "test.sqlite3")
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		err = fmt.Errorf("%s test.sqlite3", err)
		return cli.NewExitError(err, 1)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations",
		"sqlite3", driver)
	if err != nil {
		return cli.NewExitError(err, 2)
	}

	err = m.Up()
	return cli.NewExitError(err, 2)
}
