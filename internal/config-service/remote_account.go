package configservice

func (s *Service) RemoteAccountBaseURL() string {
	return s.remoteAccountBaseURL
}

func (s *Service) RemoteAccountGoogleClientID() string {
	return s.remoteAccountGoogleClientID
}

func (s *Service) RemoteAccountGoogleClientSecret() string {
	return s.remoteAccountGoogleClientSecret
}

func (s *Service) RemoteAccountYandexClientID() string {
	return s.remoteAccountYandexClientID
}

func (s *Service) RemoteAccountYandexClientSecret() string {
	return s.remoteAccountYandexClientSecret
}

func (s *Service) RemoteAccountTwitchClientID() string {
	return s.remoteAccountTwitchClientID
}

func (s *Service) RemoteAccountTwitchClientSecret() string {
	return s.remoteAccountTwitchClientSecret
}
