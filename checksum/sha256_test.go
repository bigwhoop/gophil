package checksum

import (
	"testing"
)

func TestSha256(t *testing.T) {
	s := "test"
	c := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	h := Sha256(s)
	
	if (h != c) {
		t.Errorf("Hash of string %#v should be %#v, not %#v.", s, c, h)
	}
}
