package accountservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Create(ctx context.Context, ops model.CreateAccountOptions) (uint, error) {
	if err := s.v.Struct(ops); err != nil {
		return 0, err
	}
	return s.AccountStorage.Create(ctx, ops)
}
