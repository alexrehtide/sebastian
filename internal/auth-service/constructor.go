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

type SessionService interface {
	CreateWithAccountID(ctx context.Context, accountID uint) (uint, error)
	Verify(session model.Session) error
	ReadByID(ctx context.Context, id uint) (model.Session, error)
	RefreshSession(ctx context.Context, refreshToken string) (model.Session, error)
}

func New(
	accountService AccountService,
	sessionService SessionService,
	validate validator.Validate,
) *Service {
	return &Service{
		AccountService: accountService,
		SessionService: sessionService,
		v:              validate,
	}
}

type Service struct {
	AccountService AccountService
	SessionService SessionService
	v              validator.Validate
}
