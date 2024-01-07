package password

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashPassword(src string) string {
	hasher := sha1.New()
	hasher.Write([]byte(src))
	return hex.EncodeToString(hasher.Sum(nil))
}
