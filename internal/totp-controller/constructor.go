package totpcontroller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type AccountProvider interface {
	Inject(ctx context.Context) *model.Account
}

type TOTPService interface {
	Generate(acc model.Account) (string, error)
	Validate(acc model.Account, code string) bool
}

func New(accountProvider AccountProvider, totpService TOTPService) *Controller {
	return &Controller{TOTPService: totpService, AccountProvider: accountProvider}
}

type Controller struct {
	AccountProvider AccountProvider
	TOTPService     TOTPService
}
