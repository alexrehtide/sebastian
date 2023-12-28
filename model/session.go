package model

import "time"

type Session struct {
	ID           uint      `db:"id"`
	AccountID    uint      `db:"account_id"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
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
