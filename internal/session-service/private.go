package sessionservice

import (
	"context"
	"fmt"
	"math/rand"

	customerror "github.com/alexrehtide/sebastian/internal/custom-error"
	"github.com/alexrehtide/sebastian/model"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (s *Service) generateToken() string {
	b := make([]byte, 64)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *Service) readByRefreshToken(ctx context.Context, refreshToken string) (model.Session, error) {
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
