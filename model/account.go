package model

type Account struct {
	ID       uint   `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"password_hash"`
}

type ReadAccountOptions struct {
	ID    uint
	Email string
}

type CreateAccountOptions struct {
	Email    string
	Password string
}

type UpdateAccountOptions struct {
	Email string
}
