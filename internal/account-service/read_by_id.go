package accountservice

import (
	"context"
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) ReadByID(ctx context.Context, id uint) (model.Account, error) {
	accs, err := s.AccountStorage.Read(ctx, model.ReadAccountOptions{ID: id}, model.PaginationOptions{Limit: 1})
	if err != nil {
		return model.Account{}, err
	}
	if len(accs) == 0 {
		return model.Account{}, fmt.Errorf("accountservice.Service.ReadByID: %w", serviceerror.ErrRecordNotFound)
	}
	return accs[0], nil
}
