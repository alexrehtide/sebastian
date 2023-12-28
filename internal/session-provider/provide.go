package sessionprovider

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (p *Provider) Provide(ctx context.Context, s model.Session) context.Context {
	return context.WithValue(ctx, SESSION_INJECT_KEY, s)
}
