package model

type Account struct {
	ID         uint   `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	Username   string `json:"username" db:"username"`
	Password   string `json:"-" db:"password"`
	TOTPSecret []byte `json:"-" db:"totp_secret"`
}

type ReadAccountOptions struct {
	ID       uint
	Email    string
	Username string
}

type CreateAccountOptions struct {
	Email    string `validate:"required,email,max=256"`
	Username string `validate:"required,max=256"`
	Password string `validate:"required,min=8,max=256"`
}

type UpdateAccountOptions struct {
	Username string `validate:"max=256"`
	Email    string `validate:"email,max=256"`
	Password string `validate:"required,min=8,max=256"`
}
