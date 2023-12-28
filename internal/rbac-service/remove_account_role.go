package rbacservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) RemoveAccountRole(ctx context.Context, accountID uint, role model.Role) error {
	return nil
}
