package sessionservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Verify(ctx context.Context, accessToken string) (model.Session, error) {
	session, err := s.readByAccessToken(ctx, accessToken)
	if err != nil {
		return model.Session{}, err
	}
	// TODO: verify
	return session, nil
}
