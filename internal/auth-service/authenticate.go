package authservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *AuthService) Authenticate(ctx context.Context, ops model.AuthenticateOptions) (model.Tokens, error) {
	if err := s.v.Struct(ops); err != nil {
		return model.Tokens{}, err
	}
	acc, err := s.AccountService.ReadByEmail(ctx, ops.Email)
	if err != nil {
		return model.Tokens{}, err
	}
	if err := s.verifyPassword(acc.Password, ops.Password); err != nil {
		return model.Tokens{}, err
	}
	sessionID, err := s.SessionService.CreateWithAccountID(ctx, acc.ID)
	if err != nil {
		return model.Tokens{}, err
	}
	session, err := s.SessionService.ReadByID(ctx, sessionID)
	if err != nil {
		return model.Tokens{}, err
	}
	return model.Tokens{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
	}, nil
}
