package accountcontroller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type AccountService interface {
	Count(ctx context.Context, ops model.ReadAccountOptions) (int, error)
	Create(ctx context.Context, ops model.CreateAccountOptions) (uint, error)
	Delete(ctx context.Context, id uint) error
	Read(ctx context.Context, ops model.ReadAccountOptions, pgOps model.PaginationOptions) ([]model.Account, error)
	ReadByID(ctx context.Context, id uint) (model.Account, error)
	Update(ctx context.Context, id uint, ops model.UpdateAccountOptions) error
}

func New(accountService AccountService) *Controller {
	return &Controller{
		AccountService: accountService,
	}
}

type Controller struct {
	AccountService AccountService
}
