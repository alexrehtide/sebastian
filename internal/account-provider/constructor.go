package accountprovider

const (
	ACCOUNT_INJECT_KEY = "account"
)

func New() *Provider {
	return &Provider{}
}

type Provider struct{}
