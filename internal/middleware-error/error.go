package middlewareerror

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
)
