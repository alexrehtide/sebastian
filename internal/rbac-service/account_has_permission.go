package rbacservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) AccountHasPermission(ctx context.Context, accountID uint, permission model.Permission) (bool, error) {
	return false, nil
}
