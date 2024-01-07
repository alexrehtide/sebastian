package migrator

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func New(ops MigratorOptions, fs embed.FS) (*Migrator, error) {
	dbDriver, err := postgres.WithInstance(ops.DBConn, &postgres.Config{MigrationsTable: ops.MigrationsTable})
	if err != nil {
		return nil, err
	}
	sourceDriver, err := iofs.New(fs, ".")
	if err != nil {
		return nil, err
	}
	migrate, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return nil, err
	}
	return &Migrator{
		migrate: migrate,
	}, nil
}

type MigratorOptions struct {
	DBConn          *sql.DB
	MigrationsTable string
}

type Migrator struct {
	migrate *migrate.Migrate
}

func (m *Migrator) Up() error {
	return m.migrate.Up()
}
