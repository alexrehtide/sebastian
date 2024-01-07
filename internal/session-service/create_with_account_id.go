package sessionservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/random"
)

func (s *Service) CreateWithAccountID(ctx context.Context, accountID uint) (uint, error) {
	return s.Create(ctx, model.CreateSessionOptions{
		AccountID:    accountID,
		AccessToken:  random.String(64),
		RefreshToken: random.String(64),
	})
}
