package rbacmiddleware

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type AccountProvider interface {
	Inject(ctx context.Context) *model.Account
}

type RBACService interface {
	AccountHasPermission(ctx context.Context, accountID uint, permission model.Permission) (bool, error)
}

func New(accountProvider AccountProvider, rbacService RBACService) *Middleware {
	return &Middleware{
		AccountProvider: accountProvider,
		RBACService:     rbacService,
	}
}

type Middleware struct {
	AccountProvider AccountProvider
	RBACService     RBACService
}
