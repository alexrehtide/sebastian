package authservice

import customerror "github.com/alexrehtide/sebastian/internal/custom-error"

func (s *AuthService) verifyPassword(passwordHash, password string) error {
	if passwordHash != password {
		return customerror.ErrInvalidPassword
	}
	return nil
}
