package sessionservice

import (
	"context"
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

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
		return model.Session{}, fmt.Errorf("sessionservice.Service.ReadByID: %w", serviceerror.ErrRecordNotFound)
	}
	return sessions[0], nil
}
