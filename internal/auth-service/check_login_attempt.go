package authservice

import (
	"context"
	"fmt"
	"time"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) CheckLoginAttempt(ctx context.Context, ip string) error {
	rows, err := s.LoginAttemptStorage.Read(ctx, model.ReadLoginAttemptOptions{IP: ip}, model.PaginationOptions{Limit: 1})
	if err != nil {
		return fmt.Errorf("authservice.Service.CheckLoginAttempt: %w", err)
	}
	if len(rows) != 0 {
		row := rows[0]

		if row.Count >= MAX_ATTEMPTS && time.Since(row.LastFailed) < BLOCK_DURATION {
			return fmt.Errorf("authservice.Service.CheckLoginAttempt: %w", serviceerror.ErrMaxFailedLoginAttempts)
		}
	}
	return nil
}
