package accountservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Create(ctx context.Context, ops model.CreateAccountOptions) (uint, error) {
	if err := s.v.Struct(ops); err != nil {
		return 0, fmt.Errorf("accountservice.Service.Create: %w", err)
	}
	id, err := s.AccountStorage.Create(ctx, ops)
	if err != nil {
		return 0, fmt.Errorf("accountservice.Service.Create: %w", err)
	}
	return id, nil
}
