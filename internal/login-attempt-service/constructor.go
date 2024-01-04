package loginattemptservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type LoginAttemptStorage interface {
	Create(ctx context.Context, ops model.CreateLoginAttemptOptions) (uint, error)
	Delete(ctx context.Context, ip string) error
	Read(ctx context.Context, ops model.ReadLoginAttemptOptions, pgOps model.PaginationOptions) ([]model.LoginAttempt, error)
	Update(ctx context.Context, id uint, ops model.UpdateLoginAttemptOptions) error
}

func New(loginAttemptStorage LoginAttemptStorage) *Service {
	return &Service{
		LoginAttemptStorage: loginAttemptStorage,
	}
}

type Service struct {
	LoginAttemptStorage LoginAttemptStorage
}
