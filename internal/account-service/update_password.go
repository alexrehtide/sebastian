package accountservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) UpdatePassword(ctx context.Context, accountID uint, password string) error {
	err := s.AccountStorage.Update(ctx, accountID, model.UpdateAccountOptions{
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("accountservice.Service.UpdatePassword: %w", err)
	}
	return nil
}
