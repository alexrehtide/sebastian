package loginattemptservice

import (
	"context"
	"fmt"
	"time"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) FailLoginAttempt(ctx context.Context, ip string) error {
	rows, err := s.LoginAttemptStorage.Read(ctx, model.ReadLoginAttemptOptions{IP: ip}, model.PaginationOptions{Limit: 1})
	if err != nil {
		return fmt.Errorf("authservice.AuthService.FailLoginAttempt: %w", err)
	}
	if len(rows) == 0 {
		_, err := s.LoginAttemptStorage.Create(ctx, model.CreateLoginAttemptOptions{IP: ip, Count: 1, LastFailed: time.Now()})
		if err != nil {
			return fmt.Errorf("authservice.AuthService.FailLoginAttempt: %w", err)
		}
	} else {
		row := rows[0]
		err := s.LoginAttemptStorage.Update(ctx, row.ID, model.UpdateLoginAttemptOptions{Count: row.Count + 1, LastFailed: time.Now()})
		if err != nil {
			return fmt.Errorf("authservice.AuthService.FailLoginAttempt: %w", err)
		}
	}
	return nil
}
