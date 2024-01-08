package sessionservice

import (
	"context"
	"fmt"
	"time"
)

func (s *Service) Clean(ctx context.Context) error {
	err := s.SessionStorage.DeleteOld(ctx, time.Now().Add(-s.ConfigService.SessionRefreshTokenExpiring()))
	if err != nil {
		return fmt.Errorf("sessionservice.Service.Clean: %w", err)
	}
	return nil
}
