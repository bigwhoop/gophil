package checksum

import (
	"crypto/sha256"
)

func Sha256(s string) string {
	return Generate(sha256.New(), s)
}
