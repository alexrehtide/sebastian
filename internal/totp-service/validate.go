package totpservice

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/pquerna/otp/totp"
)

func (s *Service) Validate(acc model.Account, code string) bool {
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "Sebastian",
		AccountName: acc.Email,
		Secret:      acc.TOTPSecret,
	})

	return totp.Validate(code, key.Secret())
}
