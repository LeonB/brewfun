package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/sqlite3"
	_ "github.com/mattes/migrate/source/file"
	"github.com/urfave/cli"
)

const (
	databaseName = "sqlite3"
)

func Migrate() *cli.ExitError {
	m, exitErr := getMigrate()
	defer m.Close()
	if exitErr != nil {
		return exitErr
	}

	// run all pending migrations
	err := m.Up()
	if err == nil {
		return nil
	}
	if err == migrate.ErrNoChange {
		return nil
	}
	return cli.NewExitError(err, 3)
}

func Rollback() *cli.ExitError {
	m, exitErr := getMigrate()
	defer m.Close()
	if exitErr != nil {
		return exitErr
	}

	// check version: if version == 0 == no migrations: do nothing
	version, _, _ := m.Version()
	if version == 0 {
		return nil
	}

	// roll back one step
	err := m.Steps(-1)
	if err == nil {
		return nil
	}

	// no change: don't report error
	if err == migrate.ErrNoChange {
		return nil
	}

	return cli.NewExitError(err, 5)
}

func Drop() *cli.ExitError {
	m, exitErr := getMigrate()
	defer m.Close()
	if exitErr != nil {
		return exitErr
	}

	// check version: if version == 0 == no migrations: do nothing
	version, _, _ := m.Version()
	if version == 0 {
		return nil
	}

	// roll back every migration
	err := m.Drop()
	if err == nil {
		return nil
	}

	return cli.NewExitError(err, 6)
}

func Reset() *cli.ExitError {
	exitErr := Drop()
	if exitErr != nil {
		return exitErr
	}

	return Migrate()
}

func getMigrate() (*migrate.Migrate, *cli.ExitError) {
	// m, err := migrate.New(
	// 	"file://../db/migrations",
	// 	"sqlite3://test.sqlite3?_foreign_keys=1")

	db, err := sql.Open(databaseName, "test.sqlite3?_foreign_keys=1")
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		err = fmt.Errorf("%s test.sqlite3", err)
		return nil, cli.NewExitError(err, 1)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations",
		databaseName, driver)
	if err != nil {
		return m, cli.NewExitError(err, 2)
	}

	log := &logger{}
	m.Log = log

	return m, nil
}

type logger struct {
}

func (l *logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

func (l *logger) Verbose() bool {
	return true
}
