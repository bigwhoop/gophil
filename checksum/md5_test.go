package checksum

import (
	"testing"
)

func TestMd5(t *testing.T) {
	s := "test"
	c := "098f6bcd4621d373cade4e832627b4f6"
	h := Md5(s)
	
	if (h != c) {
		t.Errorf("Hash of string %#v should be %#v, not %#v.", s, c, h)
	}
}
