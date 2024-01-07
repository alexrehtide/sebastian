package mailservice

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func New(email string, password string) *Service {
	d := gomail.NewDialer("mail.taris.fun", 465, email, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Service{
		From:   fmt.Sprintf("Taris.fun <%s>", email),
		Dialer: d,
	}
}

type Service struct {
	From   string
	Dialer *gomail.Dialer
}
