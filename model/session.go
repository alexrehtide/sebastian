package model

import "time"

type Session struct {
	ID           uint
	AccountID    uint
	AccessToken  string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ReadSessionOptions struct {
	ID           uint
	AccountID    uint
	AccessToken  string
	RefreshToken string
}

type CreateSessionOptions struct {
	AccountID    uint   `validate:"required"`
	AccessToken  string `validate:"required,len=64"`
	RefreshToken string `validate:"required,len=64"`
}

type UpdateSessionOptions struct {
	AccessToken  string `validate:"required,len=64"`
	RefreshToken string `validate:"required,len=64"`
}
