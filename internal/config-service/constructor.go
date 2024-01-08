package configservice

import "time"

func New() *Service {
	return &Service{}
}

type Service struct {
	debug          bool
	httpServerAddr string

	postgresDBName   string
	postgresUser     string
	postgresPassword string
	postgresHost     string
	postgresPort     int

	smtpHost     string
	smtpPort     int
	smtpEmail    string
	smtpPassword string

	sessionAccessTokenExpiring  time.Duration
	sessionRefreshTokenExpiring time.Duration
}
