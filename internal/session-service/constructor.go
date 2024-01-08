package sessionservice

import (
	"context"
	"time"

	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
	"github.com/sirupsen/logrus"
)

type ConfigService interface {
	SessionAccessTokenExpiring() time.Duration
	SessionRefreshTokenExpiring() time.Duration
}

type SessionStorage interface {
	Count(ctx context.Context, ops model.ReadSessionOptions) (int, error)
	Create(ctx context.Context, ops model.CreateSessionOptions) (uint, error)
	Delete(ctx context.Context, id uint) error
	DeleteOld(ctx context.Context, updatedAt time.Time) error
	Read(ctx context.Context, ops model.ReadSessionOptions, pgOps model.PaginationOptions) ([]model.Session, error)
	Update(ctx context.Context, id uint, ops model.UpdateSessionOptions) error
}

func New(configService ConfigService, log *logrus.Logger, sessionStorage SessionStorage, validate validator.Validate) *Service {
	return &Service{
		ConfigService:  configService,
		SessionStorage: sessionStorage,
		log:            log,
		v:              validate,
	}
}

type Service struct {
	ConfigService ConfigService
	SessionStorage
	log *logrus.Logger
	v   validator.Validate
}
