package mailservice

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"
)

type ConfigService interface {
	Debug() bool
	SMTPHost() string
	SMTPPort() int
	SMTPEmail() string
	SMTPPassword() string
}

func New(configService ConfigService) (*Service, error) {
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
		caCert, err := os.ReadFile("/etc/ssl/certs/certificate_ca.crt")
		if err != nil {
			return nil, fmt.Errorf("mailservice.New: %w", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		d.TLSConfig = &tls.Config{
			RootCAs: caCertPool,
		}
	}
	return &Service{
		ConfigService: configService,
		Dialer:        d,
	}, nil
}

type Service struct {
	ConfigService ConfigService
	Dialer        *gomail.Dialer
}
