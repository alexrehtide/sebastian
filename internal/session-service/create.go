package sessionservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Create(ctx context.Context, ops model.CreateSessionOptions) (uint, error) {
	if err := s.v.Struct(ops); err != nil {
		return 0, err
	}
	return s.SessionStorage.Create(ctx, ops)
}
