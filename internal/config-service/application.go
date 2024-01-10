package configservice

func (s *Service) Debug() bool {
	return s.debug
}

func (s *Service) FrontendBaseURL() string {
	return s.frontendBaseURL
}
