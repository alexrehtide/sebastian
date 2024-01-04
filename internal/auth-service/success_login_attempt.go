package authservice

import (
	"context"
	"fmt"
)

func (s *Service) SuccessLoginAttempt(ctx context.Context, ip string) error {
	err := s.LoginAttemptStorage.Delete(ctx, ip)
	if err != nil {
		return fmt.Errorf("authservice.Service.SuccessLoginAttempt: %w", err)
	}
	return nil
}
