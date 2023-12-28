package sessionservice

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
	"github.com/sirupsen/logrus"
)

type SessionStorage interface {
	Count(ctx context.Context, ops model.ReadSessionOptions) (int, error)
	Create(ctx context.Context, ops model.CreateSessionOptions) (uint, error)
	Delete(ctx context.Context, id uint) error
	Read(ctx context.Context, ops model.ReadSessionOptions, pgOps model.PaginationOptions) ([]model.Session, error)
	Update(ctx context.Context, id uint, ops model.UpdateSessionOptions) error
}

func New(log *logrus.Logger, sessionStorage SessionStorage, validate validator.Validate) *Service {
	return &Service{
		SessionStorage: sessionStorage,
		log:            log,
		v:              validate,
	}
}

type Service struct {
	SessionStorage
	log *logrus.Logger
	v   validator.Validate
}
