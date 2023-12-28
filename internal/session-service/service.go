package sessionservice

import (
	"context"
	"fmt"

	customerror "github.com/alexrehtide/sebastian/internal/custom-error"
	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
)

type SessionStorage interface {
	Count(ctx context.Context, ops model.ReadSessionOptions) (int, error)
	Create(ctx context.Context, ops model.CreateSessionOptions) (uint, error)
	Delete(ctx context.Context, id uint) error
	Read(ctx context.Context, ops model.ReadSessionOptions, pgOps model.PaginationOptions) ([]model.Session, error)
	Update(ctx context.Context, id uint, ops model.UpdateSessionOptions) error
}

func New(sessionStorage SessionStorage, validate validator.Validate) *Service {
	return &Service{
		SessionStorage: sessionStorage,
		v:              validate,
	}
}

type Service struct {
	SessionStorage
	v validator.Validate
}

func (s *Service) Create(ctx context.Context, ops model.CreateSessionOptions) (uint, error) {
	if err := s.v.Struct(ops); err != nil {
		return 0, err
	}
	return s.SessionStorage.Create(ctx, ops)
}

func (s *Service) CreateWithAccountID(ctx context.Context, accountID uint) (uint, error) {
	return s.Create(ctx, model.CreateSessionOptions{
		AccountID:    accountID,
		AccessToken:  s.GenerateToken(),
		RefreshToken: s.GenerateToken(),
	})
}

func (s *Service) GenerateToken() string {
	panic("TODO: Implement")
}

func (s *Service) Verify(ctx context.Context, accessToken string) (model.Session, error) {
	session, err := s.ReadByAccessToken(ctx, accessToken)
	if err != nil {
		return model.Session{}, err
	}
	// TODO: verify
	return session, nil
}

func (s *Service) ReadByAccessToken(ctx context.Context, accessToken string) (model.Session, error) {
	sessions, err := s.Read(
		ctx,
		model.ReadSessionOptions{
			AccessToken: accessToken,
		},
		model.PaginationOptions{
			Limit: 1,
		},
	)
	if err != nil {
		return model.Session{}, err
	}
	if len(sessions) == 0 {
		return model.Session{}, fmt.Errorf("sessionservice.Service.ReadByAccessToken: %w", customerror.ErrRecordNotFound)
	}
	return sessions[0], nil
}

func (s *Service) ReadByID(ctx context.Context, id uint) (model.Session, error) {
	sessions, err := s.Read(
		ctx,
		model.ReadSessionOptions{
			ID: id,
		},
		model.PaginationOptions{
			Limit: 1,
		},
	)
	if err != nil {
		return model.Session{}, err
	}
	if len(sessions) == 0 {
		return model.Session{}, fmt.Errorf("sessionservice.Service.ReadByID: %w", customerror.ErrRecordNotFound)
	}
	return sessions[0], nil
}

func (s *Service) RefreshSession(ctx context.Context, refreshToken string) (model.Session, error) {
	session, err := s.ReadByRefreshToken(ctx, refreshToken)
	if err != nil {
		return model.Session{}, err
	}
	// TODO: refresh
	return session, nil
}

func (s *Service) ReadByRefreshToken(ctx context.Context, refreshToken string) (model.Session, error) {
	sessions, err := s.Read(
		ctx,
		model.ReadSessionOptions{
			RefreshToken: refreshToken,
		},
		model.PaginationOptions{
			Limit: 1,
		},
	)
	if err != nil {
		return model.Session{}, err
	}
	if len(sessions) == 0 {
		return model.Session{}, fmt.Errorf("sessionservice.Service.ReadByRefreshToken: %w", customerror.ErrRecordNotFound)
	}
	return sessions[0], nil
}

func (s *Service) Update(ctx context.Context, id uint, ops model.UpdateSessionOptions) error {
	if err := s.v.Struct(ops); err != nil {
		return err
	}
	return s.SessionStorage.Update(ctx, id, ops)
}
