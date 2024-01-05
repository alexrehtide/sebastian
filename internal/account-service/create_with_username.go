package accountservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
	"github.com/sirupsen/logrus"
)

func (s *Service) CreateWithUsername(ctx context.Context, username string) (uint, error) {
	id, err := s.AccountStorage.Create(ctx, model.CreateAccountOptions{Email: "", Username: username, Password: ""})
	if err != nil {
		return 0, fmt.Errorf("accountservice.Service.CreateEmpty: %w", err)
	}
	s.log.
		WithContext(ctx).
		WithFields(logrus.Fields{
			"accountId": id,
		}).
		Info("Empty account created")
	return id, nil
}
