package authmiddleware

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type AccountProvider interface {
	Provide(ctx context.Context, acc model.Account) context.Context
}

type AccountService interface {
	ReadByID(ctx context.Context, id uint) (model.Account, error)
}

type SessionProvider interface {
	Provide(ctx context.Context, s model.Session) context.Context
}

type SessionService interface {
	ReadByAccessToken(ctx context.Context, token string) (model.Session, error)
	Verify(model.Session) error
}

func New(
	accountProvider AccountProvider,
	accountService AccountService,
	sessionProvider SessionProvider,
	sessionService SessionService,
) *Middleware {
	return &Middleware{
		AccountProvider: accountProvider,
		AccountService:  accountService,
		SessionProvider: sessionProvider,
		SessionService:  sessionService,
	}
}

type Middleware struct {
	AccountProvider AccountProvider
	AccountService  AccountService
	SessionProvider SessionProvider
	SessionService  SessionService
}
