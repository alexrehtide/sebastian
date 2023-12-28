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
	Email    string `validate:"required,email,max=256"`
	Password string `validate:"required,min=8,max=256"`
}

type UpdateAccountOptions struct {
	Email string `validate:"email,max=256"`
}
