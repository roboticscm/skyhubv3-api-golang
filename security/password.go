package security

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func EncodeSHA1Password(password string) string {
	h := sha1.New()
	io.WriteString(h, "Skyhub@010116"+password)

	return fmt.Sprintf("%x", h.Sum(nil))
}
