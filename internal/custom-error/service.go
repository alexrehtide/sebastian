package customerror

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
)
