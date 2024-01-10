package remoteaccountservice

import (
	"context"
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Exchange(ctx context.Context, platform model.Platform, code string) (string, error) {
	config, ok := s.configs[platform]
	if !ok {
		return "", fmt.Errorf("remoteaccountservice.Service.Exchange: %w", serviceerror.ErrPlatformNotFound)
	}
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return "", fmt.Errorf("remoteaccountservice.Service.Exchange: %w", err)
	}
	return token.AccessToken, nil
}
