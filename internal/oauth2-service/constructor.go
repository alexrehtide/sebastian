package oauth2service

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type RemoteAccountStorage interface {
	Create(ctx context.Context, ops model.CreateRemoteAccountOptions) (id uint, err error)
	Read(ctx context.Context, ops model.ReadRemoteAccountOptions, pgOps model.PaginationOptions) (rows []model.RemoteAccount, err error)
	Update(ctx context.Context, id uint, ops model.UpdateRemoteAccountOptions) error
}

func New(remoteAccountStorage RemoteAccountStorage, state string) *Service {
	return &Service{
		RemoteAccountStorage: remoteAccountStorage,
		State:                state,
	}
}

type Service struct {
	RemoteAccountStorage RemoteAccountStorage
	State                string
}
