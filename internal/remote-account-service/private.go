package remoteaccountservice

import (
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
	"golang.org/x/oauth2"
)

func (s *Service) getConfigByPlatform(platform model.Platform) (*oauth2.Config, error) {
	switch platform {
	case model.Google:
		return googleOauth2Config, nil
	case model.Yandex:
		return yandexOauth2Config, nil
	case model.Twitch:
		return twitchOauth2Config, nil
	}
	return nil, fmt.Errorf("remoteaccountservice.Service.getConfigByPlatform: %w", serviceerror.ErrPlatformNotFound)
}
