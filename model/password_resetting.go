package model

import "time"

type BeginPasswordResettingOptions struct {
	Email string
}

type EndPasswordResettingOptions struct {
	ResettingCode string
	NewPassword   string
}

type PasswordResetting struct {
	ID            uint      `db:"id"`
	AccountID     uint      `db:"account_id"`
	ResettingCode string    `db:"resetting_code"`
	CreatedAt     time.Time `db:"created_at"`
}

type CreatePasswordResettingOptions struct {
	AccountID     uint
	ResettingCode string
}

type ReadPasswordResettingOptions struct {
	ID            uint
	AccountID     uint
	ResettingCode string
}
