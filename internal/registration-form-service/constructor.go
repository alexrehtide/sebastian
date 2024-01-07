package registrationformservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type AccountService interface {
	Create(ctx context.Context, ops model.CreateAccountOptions) (uint, error)
}

type RBACService interface {
	AddAccountRole(ctx context.Context, accountID uint, role model.Role) error
}

type RegistrationFormStorage interface {
	Count(ctx context.Context, ops model.ReadRegistrationFormOptions) (count int, err error)
	Create(ctx context.Context, ops model.CreateRegistrationFormOptions) (id uint, err error)
	Delete(ctx context.Context, id uint) error
	Read(ctx context.Context, ops model.ReadRegistrationFormOptions, pgOps model.PaginationOptions) (rows []model.RegistrationForm, err error)
}

func New(accountService AccountService, rbacService RBACService, registrationsFormStorage RegistrationFormStorage) *Service {
	return &Service{
		AccountService:          accountService,
		RBACService:             rbacService,
		RegistrationFormStorage: registrationsFormStorage,
	}
}

type Service struct {
	AccountService          AccountService
	RBACService             RBACService
	RegistrationFormStorage RegistrationFormStorage
}
