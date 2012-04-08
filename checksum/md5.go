package checksum

import (
	"crypto/md5"
)

func Md5(s string) string {
	return Generate(md5.New(), s)
}
