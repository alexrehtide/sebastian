package model

import "time"

type BeginRegistrationOptions struct {
	Email    string
	Username string
	Password string
}

type EndRegistrationOptions struct {
	VerificationCode string
}

type Registration struct {
	ID               uint      `db:"id"`
	Email            string    `db:"email"`
	Username         string    `db:"username"`
	Password         string    `db:"password"`
	CreatedAt        time.Time `db:"created_at"`
	VerificationCode string    `db:"verification_code"`
}

type CreateRegistrationOptions struct {
	Email            string
	Username         string
	Password         string
	VerificationCode string
}

type ReadRegistrationOptions struct {
	ID               uint
	Email            string
	Username         string
	VerificationCode string
}
