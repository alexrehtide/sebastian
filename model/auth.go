package model

type AuthenticateOptions struct {
	Email    string `validate:"required,email,max=256"`
	Password string `validate:"required,min=8,max=256"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
