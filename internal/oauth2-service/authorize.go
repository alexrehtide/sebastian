package oauth2service

import (
	"context"
	"fmt"

	serviceerror "github.com/alexrehtide/sebastian/internal/service-error"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Service) Authorize(ctx context.Context, platform model.Platform, token string) (model.RemoteAccount, error) {
	var remoteAcc model.RemoteAccount
	var err error
	switch platform {
	case model.Google:
		remoteAcc, err = AuthorizeGoogle(token)
	case model.Twitch:
		remoteAcc, err = AuthorizeTwitch(token)
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
	}
	return remoteAcc, nil
}

// TODO: implement AuthorizeXXX

func AuthorizeGoogle(token string) (model.RemoteAccount, error) {
	return model.RemoteAccount{}, nil
}

func AuthorizeTwitch(token string) (model.RemoteAccount, error) {
	return model.RemoteAccount{}, nil
}

func AuthorizeYandex(token string) (model.RemoteAccount, error) {
	return model.RemoteAccount{}, nil
}
