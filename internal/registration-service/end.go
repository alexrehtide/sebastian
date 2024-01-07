package registrationservice

import (
	"context"
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) End(ctx context.Context, ops model.EndRegistrationOptions) (uint, error) {
	rows, err := s.RegistrationFormStorage.Read(ctx, model.ReadRegistrationOptions{VerificationCode: ops.VerificationCode}, model.PaginationOptions{Limit: 1})
	if err != nil {
		return 0, fmt.Errorf("registrationservice.Service.End: %w", err)
	}
	if len(rows) == 0 {
		return 0, fmt.Errorf("registrationservice.Service.End: %w", serviceerror.ErrRecordNotFound)
	}

	row := rows[0]
	accId, err := s.AccountService.Create(ctx, model.CreateAccountOptions{
		Email:    row.Email,
		Username: row.Username,
		Password: row.Password,
	})
	if err != nil {
		return 0, fmt.Errorf("registrationservice.Service.End: %w", err)
	}

	err = s.RBACService.AddAccountRole(ctx, accId, model.User)
	if err != nil {
		return 0, fmt.Errorf("registrationservice.Service.End: %w", err)
	}

	err = s.RegistrationFormStorage.Delete(ctx, row.ID)
	if err != nil {
		return 0, fmt.Errorf("registrationservice.Service.End: %w", err)
	}

	return accId, nil
}
