package authservice

import (
	"context"
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/password"
)

func (s *Service) Authenticate(ctx context.Context, ops model.AuthenticateOptions) (model.Tokens, error) {
	if err := s.v.Struct(ops); err != nil {
		return model.Tokens{}, fmt.Errorf("authservice.AuthService.Authenticate: %w", err)
	}
	acc, err := s.AccountService.ReadByEmail(ctx, ops.Email)
	if err != nil {
		return model.Tokens{}, fmt.Errorf("authservice.AuthService.Authenticate: %w", err)
	}
	if acc.Password != password.HashPassword(ops.Password) {
		return model.Tokens{}, fmt.Errorf("authservice.AuthService.Authenticate: %w", serviceerror.ErrInvalidPassword)
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
