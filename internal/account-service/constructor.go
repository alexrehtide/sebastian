package accountservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
)

type AccountStorage interface {
	Count(ctx context.Context, ops model.ReadAccountOptions) (int, error)
	Create(ctx context.Context, ops model.CreateAccountOptions) (uint, error)
	Read(ctx context.Context, ops model.ReadAccountOptions, pgOps model.PaginationOptions) ([]model.Account, error)
	Update(ctx context.Context, id uint, ops model.UpdateAccountOptions) error
	Delete(ctx context.Context, id uint) error
}

func New(accountStorage AccountStorage, validate validator.Validate) *Service {
	return &Service{
		AccountStorage: accountStorage,
		v:              validate,
	}
}

type Service struct {
	AccountStorage
	v validator.Validate
}
