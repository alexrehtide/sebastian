package registrationservice

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/random"
)

func (s *Service) Begin(ctx context.Context, ops model.BeginRegistrationOptions) (string, error) {
	code := random.String(64)
	_, err := s.RegistrationFormStorage.Create(ctx, model.CreateRegistrationOptions{
		Email:            ops.Email,
		Username:         ops.Username,
		Password:         ops.Password,
		VerificationCode: code,
	})
	if err != nil {
		return "", fmt.Errorf("registrationservice.Service.Begin: %w", err)
	}
	return fmt.Sprintf("%s/auth/sign_up?verification_code=%s", s.ConfigService.FrontendBaseURL(), code), nil // implement URL generator
}
