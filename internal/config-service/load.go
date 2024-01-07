package configservice

import (
	"fmt"
)

func (s *Service) Load() error {
	s.postgresUser = getEnv("POSTGRES_USER", "postgres")
	s.postgresPassword = getEnv("POSTGRES_PASSWORD", "3769")
	s.postgresDBName = getEnv("POSTGRES_NAME", "postgres")
	s.postgresHost = getEnv("POSTGRES_HOST", "postgres")
	postgresPort, err := getIntEnv("POSTGRES_PORT", 5432)
	if err != nil {
		return fmt.Errorf("configservice.Service.Load: %w", err)
	}
	s.postgresPort = postgresPort

	s.httpServerAddr = getEnv("HTTP_SERVER_ADDR", ":3000")
	return nil
}
