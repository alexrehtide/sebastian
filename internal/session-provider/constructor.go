package sessionprovider

const (
	SESSION_INJECT_KEY = "session"
)

func New() *Provider {
	return &Provider{}
}

type Provider struct{}
