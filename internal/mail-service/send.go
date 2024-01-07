package mailservice

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func (s *Service) Send(to string, subject string, contentType string, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody(contentType, content)
	if err := s.Dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("mailservice.Service.Send: %w", err)
	}
	return nil
}
