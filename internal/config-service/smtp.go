package configservice

func (s *Service) SMTPHost() string {
	return s.smtpHost
}

func (s *Service) SMTPPort() int {
	return s.smtpPort
}

func (s *Service) SMTPEmail() string {
	return s.smtpEmail
}

func (s *Service) SMTPPassword() string {
	return s.smtpPassword
}
