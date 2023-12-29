package accountservice

import (
	"crypto/sha1"
	"encoding/hex"
)

func (s *Service) hash(source string) string {
	hasher := sha1.New()
	hasher.Write([]byte(source))
	return hex.EncodeToString(hasher.Sum(nil))
}
