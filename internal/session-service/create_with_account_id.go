package sessionservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) CreateWithAccountID(ctx context.Context, accountID uint) (uint, error) {
	return s.Create(ctx, model.CreateSessionOptions{
		AccountID:    accountID,
		AccessToken:  s.generateToken(),
		RefreshToken: s.generateToken(),
	})
}
