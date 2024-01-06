package oauth2controller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type AccountService interface {
	CreateWithUsername(ctx context.Context, username string) (uint, error)
	ReadByID(ctx context.Context, id uint) (model.Account, error)
}

type RBACService interface {
	AddAccountRole(ctx context.Context, accountID uint, role model.Role) error
}

type SessionService interface {
	CreateWithAccountID(ctx context.Context, accountID uint) (uint, error)
	ReadByID(ctx context.Context, id uint) (model.Session, error)
}

type Oauth2Service interface {
	AuthCodeURL(platform model.Platform) (string, error)
	Authorize(ctx context.Context, platform model.Platform, token string) (model.RemoteAccount, error)
	Exchange(ctx context.Context, platform model.Platform, code string) (string, error)
	UpdateAccountID(ctx context.Context, id uint, accountID uint) error
}

func New(accountService AccountService, oauth2Service Oauth2Service, rbacService RBACService, sessionService SessionService) *Controller {
	return &Controller{
		AccountService: accountService,
		Oauth2Service:  oauth2Service,
		RBACService:    rbacService,
		SessionService: sessionService,
	}
}

type Controller struct {
	AccountService AccountService
	Oauth2Service  Oauth2Service
	RBACService    RBACService
	SessionService SessionService
}
