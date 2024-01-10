package configservice

import (
	"fmt"
	"time"

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

	s.sessionAccessTokenExpiring, err = getDurationSecondsEnv("SESSION_ACCESS_TOKEN_EXPIRING", time.Minute*15)
	if err != nil {
		return fmt.Errorf("configservice.Service.Load: %w", err)
	}
	s.sessionRefreshTokenExpiring, err = getDurationSecondsEnv("SESSION_REFRESH_TOKEN_EXPIRING", time.Hour*24)
	if err != nil {
		return fmt.Errorf("configservice.Service.Load: %w", err)
	}

	s.remoteAccountBaseURL = getEnv("REMOTE_ACCOUNT_BASE_URL", "http://localhost:9000/auth/code")
	s.remoteAccountGoogleClientID = getEnv("REMOTE_ACCOUNT_GOOGLE_CLIENT_ID", "237567719849-o5f79usjn8n49kjbk4m9fv4kig04s0fl.apps.googleusercontent.com")
	s.remoteAccountGoogleClientSecret = getEnv("REMOTE_ACCOUNT_GOOGLE_CLIENT_SECRET", "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")
	s.remoteAccountYandexClientID = getEnv("REMOTE_ACCOUNT_YANDEX_CLIENT_ID", "4dbafe42e4f04430824e9d8a6d84e68d")
	s.remoteAccountYandexClientSecret = getEnv("REMOTE_ACCOUNT_YANDEX_CLIENT_SECRET", "f15f546861b84d1abd9ed0eae3ea14dd")
	s.remoteAccountTwitchClientID = getEnv("REMOTE_ACCOUNT_TWITCH_CLIENT_ID", "imjoly4a42a44gjafzpe7nspql0wq9")
	s.remoteAccountTwitchClientSecret = getEnv("REMOTE_ACCOUNT_TWITCH_CLIENT_SECRET", "hgcwsg6o34emjpfufpvvs9hyt1xlpy")
	return nil
}
