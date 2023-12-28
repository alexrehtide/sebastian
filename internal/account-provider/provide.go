package accountprovider

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (p *Provider) Provide(ctx context.Context, acc model.Account) context.Context {
	return context.WithValue(ctx, ACCOUNT_INJECT_KEY, acc)
}
