package totpservice

import (
	"bytes"
	"fmt"
	"image/png"

	"github.com/alexrehtide/sebastian/model"
	"github.com/pquerna/otp/totp"
)

func (s *Service) Generate(acc model.Account) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Sebastian",
		AccountName: acc.Email,
		Secret:      acc.TOTPSecret,
	})
	if err != nil {
		return "", fmt.Errorf("totpservice.Service.Generate: %w", err)
	}

	var b bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		return "", fmt.Errorf("totpservice.Service.Generate: %w", err)
	}
	if err := png.Encode(&b, img); err != nil {
		return "", fmt.Errorf("totpservice.Service.Generate: %w", err)
	}

	return key.URL(), nil
}
