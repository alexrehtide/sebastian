package rbacservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
)

type AccountRoleStorage interface {
	Count(ctx context.Context, ops model.ReadAccountRoleOptions) (int, error)
	Create(ctx context.Context, ops model.CreateAccountRoleOptions) (uint, error)
	Delete(ctx context.Context, id uint) error
	Read(ctx context.Context, ops model.ReadAccountRoleOptions, pgOps model.PaginationOptions) ([]model.AccountRole, error)
}

func New(accountRoleStorage AccountRoleStorage, validate validator.Validate) *Service {
	return &Service{
		AccountRoleStorage: accountRoleStorage,
		v:                  validate,
	}
}

type Service struct {
	AccountRoleStorage
	v validator.Validate
}

func (s *Service) Create(ctx context.Context, ops model.CreateAccountRoleOptions) (uint, error) {
	if err := s.v.Struct(ops); err != nil {
		return 0, err
	}
	return s.AccountRoleStorage.Create(ctx, ops)
}

func (s *Service) AddAccountRole(ctx context.Context, accountID uint, role model.Role) error {
	return nil
}

func (s *Service) AccountHasPermission(ctx context.Context, accountID uint, permission model.Permission) (bool, error) {
	return false, nil
}

func (s *Service) ReadAccountRoles(ctx context.Context, accountID uint) ([]model.Role, error) {
	return []model.Role{}, nil
}

func (s *Service) RemoveAccountRole(ctx context.Context, accountID uint, role model.Role) error {
	return nil
}
