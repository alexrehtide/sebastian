package configservice

import "time"

func (s *Service) SessionAccessTokenExpiring() time.Duration {
	return s.sessionAccessTokenExpiring
}

func (s *Service) SessionRefreshTokenExpiring() time.Duration {
	return s.sessionRefreshTokenExpiring
}
