package controllererror

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
)
