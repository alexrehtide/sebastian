package accountservice

import (
	"context"
	"fmt"

	customerror "github.com/alexrehtide/sebastian/internal/custom-error"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) ReadByEmail(ctx context.Context, email string) (model.Account, error) {
	accs, err := s.AccountStorage.Read(ctx, model.ReadAccountOptions{Email: email}, model.PaginationOptions{Limit: 1})
	if err != nil {
		return model.Account{}, err
	}
	if len(accs) == 0 {
		return model.Account{}, fmt.Errorf("accountservice.Service.ReadByEmail: %w", customerror.ErrRecordNotFound)
	}
	return accs[0], nil
}
