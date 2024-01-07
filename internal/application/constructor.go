package application

import (
	configservice "github.com/alexrehtide/sebastian/internal/config-service"
)

type ConfigService interface {
	Load() error
	HTTPServerAddr() string
	PostgresDBName() string
	PostgresUser() string
	PostgresPassword() string
	PostgresHost() string
	PostgresPort() int
	SMTPHost() string
	SMTPPort() int
	SMTPEmail() string
	SMTPPassword() string
}

func New(configService *configservice.Service) *Application {
	return &Application{
		ConfigService: configService,
	}
}

type Application struct {
	ConfigService ConfigService
}
