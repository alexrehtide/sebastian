package passwordresettingservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type AccountService interface {
	ReadByEmail(ctx context.Context, email string) (model.Account, error)
	UpdatePassword(ctx context.Context, accountID uint, password string) error
}

type PasswordResettingStorage interface {
	Count(ctx context.Context, ops model.ReadPasswordResettingOptions) (count int, err error)
	Create(ctx context.Context, ops model.CreatePasswordResettingOptions) (id uint, err error)
	Delete(ctx context.Context, id uint) error
	Read(ctx context.Context, ops model.ReadPasswordResettingOptions, pgOps model.PaginationOptions) (rows []model.PasswordResetting, err error)
}

func New(accountService AccountService, passwordResettingStorage PasswordResettingStorage) *Service {
	return &Service{
		AccountService:           accountService,
		PasswordResettingStorage: passwordResettingStorage,
	}
}

type Service struct {
	AccountService           AccountService
	PasswordResettingStorage PasswordResettingStorage
}
