package accountservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) CreateWithUsername(ctx context.Context, username string) (uint, error) {
	id, err := s.AccountStorage.Create(ctx, model.CreateAccountOptions{Email: "", Username: username, Password: ""})
	if err != nil {
		return 0, fmt.Errorf("accountservice.Service.CreateWithUsername: %w", err)
	}
	return id, nil
}
