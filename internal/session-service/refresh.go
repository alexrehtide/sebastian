package sessionservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) RefreshSession(ctx context.Context, refreshToken string) (model.Session, error) {
	session, err := s.readByRefreshToken(ctx, refreshToken)
	if err != nil {
		return model.Session{}, err
	}
	// TODO: refresh
	return session, nil
}
