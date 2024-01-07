package configservice

func New() *Service {
	return &Service{}
}

type Service struct {
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
}
