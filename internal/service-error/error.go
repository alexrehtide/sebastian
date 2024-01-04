package serviceerror

import "errors"

var (
	ErrRecordNotFound         = errors.New("record not found")
	ErrInvalidPassword        = errors.New("invalid password")
	ErrMaxFailedLoginAttempts = errors.New("max failed login attempts")
)
