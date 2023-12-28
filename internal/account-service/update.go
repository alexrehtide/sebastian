package accountservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Update(ctx context.Context, id uint, ops model.UpdateAccountOptions) error {
	if err := s.v.Struct(ops); err != nil {
		return err
	}
	return s.AccountStorage.Update(ctx, id, ops)
}
