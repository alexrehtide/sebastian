package authservice

import "time"

const (
	MAX_ATTEMPTS   = 3
	BLOCK_DURATION = 5 * time.Minute
)
