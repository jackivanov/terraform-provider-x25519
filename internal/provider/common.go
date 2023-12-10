package provider

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

// hashForState computes the hexadecimal representation of the SHA1 checksum of a string.
func hashForState(value string) string {
	if value == "" {
		return ""
	}
	hash := sha1.Sum([]byte(strings.TrimSpace(value)))
	return hex.EncodeToString(hash[:])
}
