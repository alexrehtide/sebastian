package remoteaccountcontroller

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

type RemoteAccountService interface {
	AuthCodeURL(platform model.Platform) (string, error)
	Authorize(ctx context.Context, platform model.Platform, token string) (model.RemoteAccount, error)
	Exchange(ctx context.Context, platform model.Platform, code string) (string, error)
	UpdateAccountID(ctx context.Context, id uint, accountID uint) error
}

type SessionService interface {
	CreateWithAccountID(ctx context.Context, accountID uint) (uint, error)
	ReadByID(ctx context.Context, id uint) (model.Session, error)
}

func New(accountService AccountService, rbacService RBACService, remoteAccountService RemoteAccountService, sessionService SessionService) *Controller {
	return &Controller{
		AccountService:       accountService,
		RBACService:          rbacService,
		RemoteAccountService: remoteAccountService,
		SessionService:       sessionService,
	}
}

type Controller struct {
	AccountService       AccountService
	RBACService          RBACService
	RemoteAccountService RemoteAccountService
	SessionService       SessionService
}
