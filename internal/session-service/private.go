package sessionservice

import (
	"context"
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

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
		return model.Session{}, fmt.Errorf("sessionservice.Service.ReadByRefreshToken: %w", serviceerror.ErrRecordNotFound)
	}
	return sessions[0], nil
}
