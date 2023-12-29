package accountservice

import (
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

func (m *Service) CheckPassword(acc model.Account, password string) error {
	h := m.hash(password)
	if acc.Password != h {
		return fmt.Errorf("accountservice.Service.CheckPassword: %w", serviceerror.ErrInvalidPassword)
	}
	return nil
}
