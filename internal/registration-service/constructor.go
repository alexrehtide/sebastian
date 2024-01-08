package registrationservice

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
	Count(ctx context.Context, ops model.ReadRegistrationOptions) (count int, err error)
	Create(ctx context.Context, ops model.CreateRegistrationOptions) (id uint, err error)
	Delete(ctx context.Context, id uint) error
	Read(ctx context.Context, ops model.ReadRegistrationOptions, pgOps model.PaginationOptions) (rows []model.Registration, err error)
}

type TransactionManager interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}

func New(
	accountService AccountService,
	rbacService RBACService,
	registrationsFormStorage RegistrationFormStorage,
	trm TransactionManager,
) *Service {
	return &Service{
		AccountService:          accountService,
		RBACService:             rbacService,
		RegistrationFormStorage: registrationsFormStorage,

		trm: trm,
	}
}

type Service struct {
	AccountService          AccountService
	RBACService             RBACService
	RegistrationFormStorage RegistrationFormStorage
	trm                     TransactionManager
}
