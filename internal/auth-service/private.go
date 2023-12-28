package authservice

import (
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
)

func (s *AuthService) verifyPassword(passwordHash, password string) error {
	if passwordHash != password {
		return fmt.Errorf("authservice.AuthService.verifyPassword: %w", serviceerror.ErrInvalidPassword)
	}
	return nil
}
