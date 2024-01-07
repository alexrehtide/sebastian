package passwordresettingservice

import (
	"context"
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) End(ctx context.Context, ops model.EndPasswordResettingOptions) error {
	rows, err := s.PasswordResettingStorage.Read(ctx, model.ReadPasswordResettingOptions{ResettingCode: ops.ResettingCode}, model.PaginationOptions{Limit: 1})
	if err != nil {
		return fmt.Errorf("passwordresettingservice.Service.End: %w", err)
	}
	if len(rows) == 0 {
		return fmt.Errorf("passwordresettingservice.Service.End: %w", serviceerror.ErrRecordNotFound)
	}

	row := rows[0]
	err = s.AccountService.UpdatePassword(ctx, row.AccountID, ops.NewPassword)
	if err != nil {
		return fmt.Errorf("passwordresettingservice.Service.End: %w", err)
	}

	err = s.PasswordResettingStorage.Delete(ctx, row.ID)
	if err != nil {
		return fmt.Errorf("passwordresettingservice.Service.End: %w", err)
	}

	return nil
}
