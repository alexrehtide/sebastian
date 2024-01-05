package oauth2controller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type AccountService interface {
	CreateWithUsername(ctx context.Context, username string) (uint, error)
	ReadByID(ctx context.Context, id uint) (model.Account, error)
}

type SessionService interface {
	CreateWithAccountID(ctx context.Context, accountID uint) (uint, error)
	ReadByID(ctx context.Context, id uint) (model.Session, error)
}

type Oauth2Service interface {
	AuthCodeURL(platform model.Platform) (string, error)
	Authorize(ctx context.Context, platform model.Platform, token string) (model.RemoteAccount, error)
	Exchange(ctx context.Context, platform model.Platform, code string) (string, error)
}

func New(accountService AccountService, oauth2Service Oauth2Service, sessionService SessionService) *Controller {
	return &Controller{
		AccountService: accountService,
		Oauth2Service:  oauth2Service,
		SessionService: sessionService,
	}
}

type Controller struct {
	AccountService AccountService
	SessionService SessionService
	Oauth2Service  Oauth2Service
}
