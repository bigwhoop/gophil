# Some utils for the Go programming language

## Checksum

### Usage

    import "github.com/bigwhoop/gophil/checksum"
    hash := checksum.Md5("test")
    hash := checksum.Sha256("test")

or

    import (
        "github.com/bigwhoop/gophil/checksum"
        "crypto/md5"
    )
    hash := checksum.Generate(md5.New(), "test")