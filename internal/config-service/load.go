package configservice

import (
	"fmt"

	"github.com/joho/godotenv"
)

func (s *Service) Load() (err error) {
	err = godotenv.Load()
	if err != nil {
		return fmt.Errorf("configservice.Service.Load: %w", err)
	}

	s.httpServerAddr = getEnv("HTTP_SERVER_ADDR", ":3000")
	s.debug, err = getBoolEnv("DEBUG", false)
	if err != nil {
		return fmt.Errorf("configservice.Service.Load: %w", err)
	}

	s.postgresUser = getEnv("POSTGRES_USER", "postgres")
	s.postgresPassword = getEnv("POSTGRES_PASSWORD", "3769")
	s.postgresDBName = getEnv("POSTGRES_NAME", "postgres")
	s.postgresHost = getEnv("POSTGRES_HOST", "localhost")
	s.postgresPort, err = getIntEnv("POSTGRES_PORT", 5432)
	if err != nil {
		return fmt.Errorf("configservice.Service.Load: %w", err)
	}

	s.smtpHost = getEnv("SMTP_HOST", "mail.taris.fun")
	s.smtpPort, err = getIntEnv("SMTP_PORT", 465)
	if err != nil {
		return fmt.Errorf("configservice.Service.Load: %w", err)
	}
	s.smtpEmail = getEnv("SMTP_EMAIL", "admin@taris.fun")
	s.smtpPassword = getEnv("SMTP_PASSWORD", "32213345Qq")

	return nil
}
