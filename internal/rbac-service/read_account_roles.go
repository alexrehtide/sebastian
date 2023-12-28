package rbacservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) ReadAccountRoles(ctx context.Context, accountID uint) ([]model.Role, error) {
	return []model.Role{}, nil
}
