package rbacservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Create(ctx context.Context, ops model.CreateAccountRoleOptions) (uint, error) {
	if err := s.v.Struct(ops); err != nil {
		return 0, err
	}
	return s.AccountRoleStorage.Create(ctx, ops)
}
