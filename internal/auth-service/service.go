package authservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
)

type AccountService interface {
	ReadByEmail(ctx context.Context, email string) (model.Account, error)
	ReadByID(ctx context.Context, id uint) (model.Account, error)
}

type SessionService interface {
	CreateWithAccountID(ctx context.Context, accountID uint) (uint, error)
	Verify(ctx context.Context, accessToken string) (model.Session, error)
	ReadByID(ctx context.Context, id uint) (model.Session, error)
	RefreshSession(ctx context.Context, refreshToken string) (model.Session, error)
}

func New(accountService AccountService, sessionService SessionService, validate validator.Validate) *AuthService {
	return &AuthService{
		AccountService: accountService,
		SessionService: sessionService,
		v:              validate,
	}
}

type AuthService struct {
	AccountService AccountService
	SessionService SessionService
	v              validator.Validate
}

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

func (s *AuthService) verifyPassword(passwordHash, password string) error {
	panic("TODO: Implement")
}
