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
