package oauth2service

import (
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) AuthCodeURL(platform model.Platform) (string, error) {
	config, err := s.getConfigByPlatform(platform)
	if err != nil {
		return "", fmt.Errorf("oauth2service.Service.AuthCodeURL: %w", err)
	}
	return config.AuthCodeURL(s.State), nil
}
