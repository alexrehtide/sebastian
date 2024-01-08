package application

import (
	"fmt"

	"github.com/alexrehtide/sebastian/migrations"
	"github.com/alexrehtide/sebastian/pkg/migrator"
)

func (a *Application) MigrateDB() error {
	err := a.ConfigService.Load()
	if err != nil {
		return fmt.Errorf("application.Application.MigrateDB: %w", err)
	}

	sqlDB, err := a.dbConnection()
	if err != nil {
		return fmt.Errorf("application.Application.MigrateDB: %w", err)
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
		return fmt.Errorf("application.Application.MigrateDB: %w", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		return fmt.Errorf("application.Application.MigrateDB: %w", err)
	}

	return nil
}
