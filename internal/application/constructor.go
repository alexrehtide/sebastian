package application

import (
	"time"

	configservice "github.com/alexrehtide/sebastian/internal/config-service"
)

type ConfigService interface {
	Load() error
	Debug() bool
	FrontendBaseURL() string
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
	SessionAccessTokenExpiring() time.Duration
	SessionRefreshTokenExpiring() time.Duration
	RemoteAccountBaseURL() string
	RemoteAccountGoogleClientID() string
	RemoteAccountGoogleClientSecret() string
	RemoteAccountYandexClientID() string
	RemoteAccountYandexClientSecret() string
	RemoteAccountTwitchClientID() string
	RemoteAccountTwitchClientSecret() string
}

func New(configService *configservice.Service) *Application {
	return &Application{
		ConfigService: configService,
	}
}

type Application struct {
	ConfigService ConfigService
}
