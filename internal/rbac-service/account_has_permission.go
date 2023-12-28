package rbacservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) AccountHasPermission(ctx context.Context, accountID uint, permission model.Permission) (bool, error) {
	roles, err := s.ReadAccountRoles(ctx, accountID)
	if err != nil {
		return false, fmt.Errorf("rbacservice.Service.AccountHasPermission: %w", err)
	}
	for _, r := range roles {
		for _, p := range model.RolePermission[r] {
			if p == permission {
				return true, nil
			}
		}
	}
	return false, nil
}
