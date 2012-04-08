package checksum

import (
	"fmt"
	"hash"
)

func Generate(h hash.Hash, s string) string {
	h.Reset()
	h.Write([]byte(s))
	data := h.Sum([]byte{})
	return fmt.Sprintf("%x", data)
}
