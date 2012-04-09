package fs

import (
	"io"
	"os"
)

// Copies a file to another directory
func CopyFile(src, dst string) (int64, error) {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0)
	if err != nil {
		return 0, err
	}

	dstFile, err := os.OpenFile(dst, os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}

	return io.Copy(dstFile, srcFile)
}
