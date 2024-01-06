package rbacservice

import (
	"context"
	"fmt"
	"slices"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) AddAccountRole(ctx context.Context, accountID uint, role model.Role) error {
	roles, err := s.ReadAccountRoles(ctx, accountID)
	if err != nil {
		return fmt.Errorf("rbacservice.Service.AddAccountRole: %w", err)
	}
	if slices.Contains(roles, role) {
		return nil
	}
	_, err = s.Create(ctx, model.CreateAccountRoleOptions{AccountID: accountID, Role: role})
	if err != nil {
		return fmt.Errorf("rbacservice.Service.AddAccountRole: %w", err)
	}
	return nil
}
