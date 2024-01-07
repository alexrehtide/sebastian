package application

import (
	"fmt"

	"github.com/alexrehtide/sebastian/migrations"
	"github.com/alexrehtide/sebastian/pkg/migrator"
	"github.com/alexrehtide/sebastian/pkg/postgres"
)

func (a *Application) MigrateDB() error {
	err := a.ConfigService.Load()
	if err != nil {
		return fmt.Errorf("httpserver.Server.MigrateDB: %w", err)
	}

	sqlDB, err := postgres.New(postgres.PostgresOptions{
		User:     a.ConfigService.PostgresUser(),
		Password: a.ConfigService.PostgresPassword(),
		Host:     a.ConfigService.PostgresHost(),
		Port:     a.ConfigService.PostgresPort(),
		DBName:   a.ConfigService.PostgresDBName(),
	})
	if err != nil {
		return fmt.Errorf("httpserver.Server.MigrateDB: %w", err)
	}
	defer sqlDB.Close()

	m, err := migrator.New(
		migrator.MigratorOptions{
			DBConn:          sqlDB,
			MigrationsTable: "migrations_table",
		},
		migrations.FS,
	)
	if err != nil {
		return fmt.Errorf("httpserver.Server.MigrateDB: %w", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		return fmt.Errorf("httpserver.Server.MigrateDB: %w", err)
	}

	return nil
}
