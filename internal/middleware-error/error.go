package middlewareerror

import "errors"

var (
	ErrSessionNotFound  = errors.New("session not found")
	ErrAccountNotFound  = errors.New("account not found")
	ErrPermissionDenied = errors.New("permission denied")
)
