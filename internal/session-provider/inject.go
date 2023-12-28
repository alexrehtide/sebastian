package sessionprovider

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (p *Provider) Inject(ctx context.Context) *model.Session {
	acc, ok := ctx.Value(SESSION_INJECT_KEY).(model.Session)
	if !ok {
		return nil
	}
	return &acc
}
