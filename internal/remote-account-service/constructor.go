package remoteaccountservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/twitch"
	"golang.org/x/oauth2/yandex"
)

type ConfigService interface {
	RemoteAccountBaseURL() string
	RemoteAccountGoogleClientID() string
	RemoteAccountGoogleClientSecret() string
	RemoteAccountYandexClientID() string
	RemoteAccountYandexClientSecret() string
	RemoteAccountTwitchClientID() string
	RemoteAccountTwitchClientSecret() string
}

type RemoteAccountStorage interface {
	Create(ctx context.Context, ops model.CreateRemoteAccountOptions) (id uint, err error)
	Read(ctx context.Context, ops model.ReadRemoteAccountOptions, pgOps model.PaginationOptions) (rows []model.RemoteAccount, err error)
	Update(ctx context.Context, id uint, ops model.UpdateRemoteAccountOptions) error
}

func New(configService ConfigService, remoteAccountStorage RemoteAccountStorage, state string) *Service {
	configs := make(map[model.Platform]*oauth2.Config)

	configs[model.Google] = &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/%s", configService.RemoteAccountBaseURL(), model.Google),
		ClientID:     configService.RemoteAccountGoogleClientID(),
		ClientSecret: configService.RemoteAccountGoogleClientSecret(),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	configs[model.Yandex] = &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/%s", configService.RemoteAccountBaseURL(), model.Yandex),
		ClientID:     configService.RemoteAccountYandexClientID(),
		ClientSecret: configService.RemoteAccountYandexClientSecret(),
		Scopes:       []string{},
		Endpoint:     yandex.Endpoint,
	}

	configs[model.Twitch] = &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/%s", configService.RemoteAccountBaseURL(), model.Twitch),
		ClientID:     configService.RemoteAccountTwitchClientID(),
		ClientSecret: configService.RemoteAccountTwitchClientSecret(),
		Scopes:       []string{"user:read:email"},
		Endpoint:     twitch.Endpoint,
	}

	return &Service{
		ConfigService:        configService,
		RemoteAccountStorage: remoteAccountStorage,
		State:                state,

		configs: configs,
	}
}

type Service struct {
	ConfigService        ConfigService
	RemoteAccountStorage RemoteAccountStorage
	State                string

	configs map[model.Platform]*oauth2.Config
}
