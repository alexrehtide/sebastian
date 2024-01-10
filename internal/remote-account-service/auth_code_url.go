package remoteaccountservice

import (
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) AuthCodeURL(platform model.Platform) (string, error) {
	config, ok := s.configs[platform]
	if !ok {
		return "", fmt.Errorf("oauth2service.Service.AuthCodeURL: %w", serviceerror.ErrPlatformNotFound)
	}
	return config.AuthCodeURL(s.State), nil
}
