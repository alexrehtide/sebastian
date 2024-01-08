package application

import (
	"database/sql"
	"fmt"

	"github.com/alexrehtide/sebastian/pkg/postgres"
)

func (a *Application) dbConnection() (*sql.DB, error) {
	sqlDB, err := postgres.New(postgres.PostgresOptions{
		User:     a.ConfigService.PostgresUser(),
		Password: a.ConfigService.PostgresPassword(),
		Host:     a.ConfigService.PostgresHost(),
		Port:     a.ConfigService.PostgresPort(),
		DBName:   a.ConfigService.PostgresDBName(),
	})
	if err != nil {
		return nil, fmt.Errorf("application.Application.dbConnection: %w", err)
	}
	return sqlDB, nil
}
