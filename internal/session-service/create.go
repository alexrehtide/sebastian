package sessionservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
	"github.com/sirupsen/logrus"
)

func (s *Service) Create(ctx context.Context, ops model.CreateSessionOptions) (uint, error) {
	if err := s.v.Struct(ops); err != nil {
		return 0, fmt.Errorf("sessionservice.Service.Create: %w", err)
	}
	id, err := s.SessionStorage.Create(ctx, ops)
	if err != nil {
		return 0, fmt.Errorf("sessionservice.Service.Create: %w", err)
	}
	s.log.WithContext(ctx).
		WithFields(logrus.Fields{
			"sessionAccessToken":  ops.AccessToken,
			"sessionRefreshToken": ops.RefreshToken,
			"sessionAccountId":    ops.AccountID,
			"sessionId":           id,
		}).
		Info("Session created")
	return id, nil
}
