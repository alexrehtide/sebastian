package authservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *AuthService) Authenticate(ctx context.Context, ops model.AuthenticateOptions) (model.Tokens, error) {
	if err := s.v.Struct(ops); err != nil {
		return model.Tokens{}, fmt.Errorf("authservice.AuthService.Authenticate: %w", err)
	}
	acc, err := s.AccountService.ReadByEmail(ctx, ops.Email)
	if err != nil {
		return model.Tokens{}, fmt.Errorf("authservice.AuthService.Authenticate: %w", err)
	}
	if err := s.AccountService.CheckPassword(acc, ops.Password); err != nil {
		return model.Tokens{}, fmt.Errorf("authservice.AuthService.Authenticate: %w", err)
	}
	sessionID, err := s.SessionService.CreateWithAccountID(ctx, acc.ID)
	if err != nil {
		return model.Tokens{}, fmt.Errorf("authservice.AuthService.Authenticate: %w", err)
	}
	session, err := s.SessionService.ReadByID(ctx, sessionID)
	if err != nil {
		return model.Tokens{}, fmt.Errorf("authservice.AuthService.Authenticate: %w", err)
	}
	return model.Tokens{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
	}, nil
}
