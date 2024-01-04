package authservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
)

type AccountService interface {
	CheckPassword(account model.Account, password string) error
	ReadByEmail(ctx context.Context, email string) (model.Account, error)
	ReadByID(ctx context.Context, id uint) (model.Account, error)
}

type LoginAttemptStorage interface {
	Create(ctx context.Context, ops model.CreateLoginAttemptOptions) (uint, error)
	Delete(ctx context.Context, ip string) error
	Read(ctx context.Context, ops model.ReadLoginAttemptOptions, pgOps model.PaginationOptions) ([]model.LoginAttempt, error)
	Update(ctx context.Context, id uint, ops model.UpdateLoginAttemptOptions) error
}

type SessionService interface {
	CreateWithAccountID(ctx context.Context, accountID uint) (uint, error)
	Verify(session model.Session) error
	ReadByID(ctx context.Context, id uint) (model.Session, error)
	RefreshSession(ctx context.Context, refreshToken string) (model.Session, error)
}

func New(
	accountService AccountService,
	loginAttemptStorage LoginAttemptStorage,
	sessionService SessionService,
	validate validator.Validate,
) *Service {
	return &Service{
		AccountService:      accountService,
		LoginAttemptStorage: loginAttemptStorage,
		SessionService:      sessionService,
		v:                   validate,
	}
}

type Service struct {
	AccountService      AccountService
	LoginAttemptStorage LoginAttemptStorage
	SessionService      SessionService
	v                   validator.Validate
}
