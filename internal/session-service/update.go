package sessionservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Update(ctx context.Context, id uint, ops model.UpdateSessionOptions) error {
	if err := s.v.Struct(ops); err != nil {
		return err
	}
	return s.SessionStorage.Update(ctx, id, ops)
}
