package oauth2service

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) UpdateAccountID(ctx context.Context, id uint, accountID uint) error {
	if err := s.RemoteAccountStorage.Update(ctx, id, model.UpdateRemoteAccountOptions{AccountID: accountID}); err != nil {
		return fmt.Errorf("oauth2service.Service.UpdateAccountID")
	}
	return nil
}
