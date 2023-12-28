package accountservice

import (
	"context"
	"fmt"

	customerror "github.com/alexrehtide/sebastian/internal/custom-error"
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

func (s *Service) Create(ctx context.Context, ops model.CreateAccountOptions) (uint, error) {
	if err := s.v.Struct(ops); err != nil {
		return 0, err
	}
	return s.AccountStorage.Create(ctx, ops)
}

func (s *Service) ReadByEmail(ctx context.Context, email string) (model.Account, error) {
	accs, err := s.AccountStorage.Read(ctx, model.ReadAccountOptions{Email: email}, model.PaginationOptions{Limit: 1})
	if err != nil {
		return model.Account{}, err
	}
	if len(accs) == 0 {
		return model.Account{}, fmt.Errorf("accountservice.Service.ReadByEmail: %w", customerror.ErrRecordNotFound)
	}
	return accs[0], nil
}

func (s *Service) ReadByID(ctx context.Context, id uint) (model.Account, error) {
	accs, err := s.AccountStorage.Read(ctx, model.ReadAccountOptions{ID: id}, model.PaginationOptions{Limit: 1})
	if err != nil {
		return model.Account{}, err
	}
	if len(accs) == 0 {
		return model.Account{}, fmt.Errorf("accountservice.Service.ReadByID: %w", customerror.ErrRecordNotFound)
	}
	return accs[0], nil
}

func (s *Service) Update(ctx context.Context, id uint, ops model.UpdateAccountOptions) error {
	if err := s.v.Struct(ops); err != nil {
		return err
	}
	return s.AccountStorage.Update(ctx, id, ops)
}
