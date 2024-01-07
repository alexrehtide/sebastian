package remoteaccountservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Exchange(ctx context.Context, platform model.Platform, code string) (string, error) {
	config, err := s.getConfigByPlatform(platform)
	if err != nil {
		return "", fmt.Errorf("remoteaccountservice.Service.Exchange: %w", err)
	}
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return "", fmt.Errorf("remoteaccountservice.Service.Exchange: %w", err)
	}
	return token.AccessToken, nil
}
