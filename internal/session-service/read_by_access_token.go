package sessionservice

import (
	"context"
	"fmt"

	customerror "github.com/alexrehtide/sebastian/internal/custom-error"
	"github.com/alexrehtide/sebastian/model"
)

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
