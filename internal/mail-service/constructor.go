package mailservice

import (
	"crypto/tls"

	gomail "gopkg.in/mail.v2"
)

type ConfigService interface {
	Debug() bool
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
	if configService.Debug() {
		d.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	} else {
		d.TLSConfig = &tls.Config{
			ServerName: "taris.fun",
		}
	}
	return &Service{
		ConfigService: configService,
		Dialer:        d,
	}
}

type Service struct {
	ConfigService ConfigService
	Dialer        *gomail.Dialer
}
