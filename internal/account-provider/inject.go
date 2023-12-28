package accountprovider

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (p *Provider) Inject(ctx context.Context) *model.Account {
	acc, ok := ctx.Value(ACCOUNT_INJECT_KEY).(model.Account)
	if !ok {
		return nil
	}
	return &acc
}
