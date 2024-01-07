package application

import (
	configservice "github.com/alexrehtide/sebastian/internal/config-service"
	"github.com/gofiber/fiber/v2"
)

type ConfigService interface {
	Load() error
	PostgresDBName() string
	PostgresUser() string
	PostgresPassword() string
	PostgresHost() string
	PostgresPort() int
	HTTPServerAddr() string
}

func New(configService *configservice.Service) *Application {
	return &Application{
		ConfigService: configService,
	}
}

type Application struct {
	ConfigService ConfigService

	App *fiber.App
}
