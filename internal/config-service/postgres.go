package configservice

func (s *Service) PostgresDBName() string {
	return s.postgresDBName
}

func (s *Service) PostgresUser() string {
	return s.postgresUser
}

func (s *Service) PostgresPassword() string {
	return s.postgresPassword
}

func (s *Service) PostgresHost() string {
	return s.postgresHost
}

func (s *Service) PostgresPort() int {
	return s.postgresPort
}
