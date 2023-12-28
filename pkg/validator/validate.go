package validator

import "github.com/go-playground/validator/v10"

type Validate interface {
	Struct(s interface{}) error
}

func New() Validate {
	return &validate{
		v: validator.New(validator.WithRequiredStructEnabled()),
	}
}

type validate struct {
	v *validator.Validate
}

func (v *validate) Struct(s interface{}) error {
	return v.v.Struct(s)
}
