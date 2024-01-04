package model

import "time"

type LoginAttempt struct {
	ID         uint      `db:"id"`
	IP         string    `db:"ip"`
	Count      int       `db:"count"`
	LastFailed time.Time `db:"last_failed"`
}

type ReadLoginAttemptOptions struct {
	ID uint
	IP string
}

type CreateLoginAttemptOptions struct {
	IP         string `validate:"required"`
	Count      int
	LastFailed time.Time `validate:"required"`
}

type UpdateLoginAttemptOptions struct {
	Count      int
	LastFailed time.Time `validate:"required"`
}
