package authservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
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

func New(accountService AccountService, sessionService SessionService) *AuthService {
	return &AuthService{
		AccountService: accountService,
		SessionService: sessionService,
	}
}

type AuthService struct {
	AccountService AccountService
	SessionService SessionService
}

func (a *AuthService) Authenticate(ctx context.Context, in model.AuthenticateInput) (model.AuthenticateOutput, error) {
	acc, err := a.AccountService.ReadByEmail(ctx, in.Email)
	if err != nil {
		return model.AuthenticateOutput{}, err
	}
	if err := a.verifyPassword(acc.Password, in.Password); err != nil {
		return model.AuthenticateOutput{}, err
	}
	sessionID, err := a.SessionService.CreateWithAccountID(ctx, acc.ID)
	if err != nil {
		return model.AuthenticateOutput{}, err
	}
	session, err := a.SessionService.ReadByID(ctx, sessionID)
	if err != nil {
		return model.AuthenticateOutput{}, err
	}
	return model.AuthenticateOutput{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
	}, nil
}

func (a *AuthService) verifyPassword(passwordHash, password string) error {
	panic("TODO: Implement")
}
