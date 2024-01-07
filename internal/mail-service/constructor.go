package mailservice

import (
	"crypto/tls"

	gomail "gopkg.in/mail.v2"
)

type ConfigService interface {
	SMTPHost() string
	SMTPPort() int
	SMTPEmail() string
	SMTPPassword() string
}

func New(configService ConfigService) *Service {
	d := gomail.NewDialer(
		configService.SMTPHost(),
		configService.SMTPPort(),
		configService.SMTPEmail(),
		configService.SMTPPassword(),
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Service{
		ConfigService: configService,
		Dialer:        d,
	}
}

type Service struct {
	ConfigService ConfigService
	Dialer        *gomail.Dialer
}
