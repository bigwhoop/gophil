# Some utils for the Go programming language

## Core

### MinInt([]int) int, MaxInt([]int) int

    import "github.com/bigwhoop/gophil/core"
    ints := []int{3,-5,12,0,28,4}
    i := core.MinInt(ints) // -5
    i := core.MaxInt(ints) // 28


## Fs

### CopyFile(src, dst string) (int64, error)

    import "github.com/bigwhoop/gophil/fs"
    bytesWritten, err := fs.CopyFile("~/source.go", "~/dest.go")


## Checksum

### Md5(string) string, Sha256(string) string

    import "github.com/bigwhoop/gophil/checksum"
    hash := checksum.Md5("example string")
    hash := checksum.Sha256("example string")

### Generate(hash.Hash, string) string

    import (
        "github.com/bigwhoop/gophil/checksum"
        "crypto/md5"
    )
    checksum.Generate(md5.New(), "example string")