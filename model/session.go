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
	AccountID    uint
	AccessToken  string
	RefreshToken string
}

type UpdateSessionOptions struct {
	AccessToken  string
	RefreshToken string
}
