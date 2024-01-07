package passwordresettingservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/random"
)

func (s *Service) Begin(ctx context.Context, ops model.BeginPasswordResettingOptions) (string, error) {
	acc, err := s.AccountService.ReadByEmail(ctx, ops.Email)
	if err != nil {
		return "", fmt.Errorf("passwordresettingservice.Service.Begin: %w", err)
	}

	code := random.String(64)
	_, err = s.PasswordResettingStorage.Create(ctx, model.CreatePasswordResettingOptions{
		AccountID:     acc.ID,
		ResettingCode: code,
	})
	if err != nil {
		return "", fmt.Errorf("passwordresettingservice.Service.Begin: %w", err)
	}
	return fmt.Sprintf("http://localhost:9000/auth/forgot_password?resetting_code=%s", code), nil // implement URL generator
}
