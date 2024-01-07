package oauth2service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
	"golang.org/x/oauth2"
)

func (s *Service) Authorize(ctx context.Context, platform model.Platform, token string) (model.RemoteAccount, error) {
	var remoteAcc model.RemoteAccount
	var err error
	switch platform {
	case model.Google:
		remoteAcc, err = AuthorizeGoogle(token)
	case model.Twitch:
		remoteAcc, err = AuthorizeTwitch(twitchOauth2Config, token)
	case model.Yandex:
		remoteAcc, err = AuthorizeYandex(token)
	default:
		return model.RemoteAccount{}, serviceerror.ErrPlatformNotFound
	}
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.Service.Authorize: %w", err)
	}
	rows, err := s.RemoteAccountStorage.Read(
		ctx,
		model.ReadRemoteAccountOptions{
			RemoteID: remoteAcc.RemoteID,
			Platform: remoteAcc.Platform,
		},
		model.PaginationOptions{
			Limit: 1,
		},
	)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.Service.Authorize: %w", err)
	}
	if len(rows) == 0 {
		id, err := s.RemoteAccountStorage.Create(
			ctx,
			model.CreateRemoteAccountOptions{
				RemoteID:    remoteAcc.RemoteID,
				RemoteEmail: remoteAcc.RemoteEmail,
				Platform:    remoteAcc.Platform,
			},
		)
		if err != nil {
			return model.RemoteAccount{}, fmt.Errorf("oauth2service.Service.Authorize: %w", err)
		}
		remoteAcc.ID = id
	} else {
		row := rows[0]
		remoteAcc.ID = row.ID
		remoteAcc.AccountID = row.AccountID
	}
	return remoteAcc, nil
}

type GoogleUserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func AuthorizeGoogle(token string) (model.RemoteAccount, error) {
	url := "https://www.googleapis.com/oauth2/v2/userinfo"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeGoogle: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeGoogle: %w", err)
	}

	var userInfo GoogleUserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeGoogle: %w", err)
	}

	return model.RemoteAccount{
		RemoteID:    userInfo.ID,
		RemoteEmail: userInfo.Email,
		Platform:    model.Google,
	}, nil
}

type TwitchUserInfo struct {
	Data []struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	} `json:"data"`
}

func AuthorizeTwitch(config *oauth2.Config, token string) (model.RemoteAccount, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users", nil)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeTwitch: %w", err)
	}
	req.Header.Add("Client-ID", config.ClientID)
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeTwitch: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeTwitch: %w", err)
	}

	var userInfo TwitchUserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeTwitch: %w", err)
	}

	if len(userInfo.Data) == 0 {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeTwitch: %w", serviceerror.ErrRecordNotFound)
	}

	return model.RemoteAccount{
		RemoteID:    userInfo.Data[0].ID,
		RemoteEmail: userInfo.Data[0].Email,
		Platform:    model.Twitch,
	}, nil
}

type YandexUserInfo struct {
	ID    string `json:"id"`
	Email string `json:"default_email"`
}

func AuthorizeYandex(token string) (model.RemoteAccount, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://login.yandex.ru/info", nil)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeYandex: %w", err)
	}
	req.Header.Add("Authorization", "OAuth "+token)

	resp, err := client.Do(req)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeYandex: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeYandex: %w", err)
	}

	var userInfo YandexUserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return model.RemoteAccount{}, fmt.Errorf("oauth2service.AuthorizeYandex: %w", err)
	}

	return model.RemoteAccount{
		RemoteID:    userInfo.ID,
		RemoteEmail: userInfo.Email,
		Platform:    model.Yandex,
	}, nil
}
