package accountservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
	"github.com/sirupsen/logrus"
)

func (s *Service) Create(ctx context.Context, ops model.CreateAccountOptions) (uint, error) {
	if err := s.v.Struct(ops); err != nil {
		return 0, fmt.Errorf("accountservice.Service.Create: %w", err)
	}
	id, err := s.AccountStorage.Create(ctx, ops)
	if err != nil {
		return 0, fmt.Errorf("accountservice.Service.Create: %w", err)
	}
	s.log.
		WithContext(ctx).
		WithFields(logrus.Fields{
			"accountId":    id,
			"accountEmail": ops.Email,
		}).
		Info("Account created")
	return id, nil
}
