package remoteaccountservice

import (
	"fmt"

	"github.com/alexrehtide/sebastian/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/twitch"
	"golang.org/x/oauth2/yandex"
)

const (
	BaseCallbackURL = "http://localhost:9000/auth/code"
)

var googleOauth2Config = &oauth2.Config{
	RedirectURL:  fmt.Sprintf("%s/%s", BaseCallbackURL, model.Google),
	ClientID:     "237567719849-o5f79usjn8n49kjbk4m9fv4kig04s0fl.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

var yandexOauth2Config = &oauth2.Config{
	RedirectURL:  fmt.Sprintf("%s/%s", BaseCallbackURL, model.Yandex),
	ClientID:     "4dbafe42e4f04430824e9d8a6d84e68d",
	ClientSecret: "f15f546861b84d1abd9ed0eae3ea14dd",
	Scopes:       []string{},
	Endpoint:     yandex.Endpoint,
}

var twitchOauth2Config = &oauth2.Config{
	RedirectURL:  fmt.Sprintf("%s/%s", BaseCallbackURL, model.Twitch),
	ClientID:     "imjoly4a42a44gjafzpe7nspql0wq9",
	ClientSecret: "hgcwsg6o34emjpfufpvvs9hyt1xlpy",
	Scopes:       []string{"user:read:email"},
	Endpoint:     twitch.Endpoint,
}
