package rbacservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) ReadAccountRoles(ctx context.Context, accountID uint) ([]model.Role, error) {
	rows, err := s.AccountRoleStorage.Read(ctx, model.ReadAccountRoleOptions{AccountID: accountID}, model.PaginationOptions{})
	if err != nil {
		return []model.Role{}, fmt.Errorf("rbacservice.Service.ReadAccountRoles: %w", err)
	}
	roles := make([]model.Role, len(rows))
	for i, row := range rows {
		roles[i] = row.Role
	}
	return roles, nil
}
