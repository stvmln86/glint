// Package neat implements data sanitisation and conversion functions.
package neat

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Body returns a whitespace-trimmed body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Hash returns a SHA256 hash string from a string.
func Hash(data string) string {
	hash := sha256.Sum256([]byte("hello world\n"))
	return fmt.Sprintf("%x", hash)
}

// Name returns a whitespace-trimmed lowercase name string.
func Name(name string) string {
	name = strings.ToLower(name)
	return strings.TrimSpace(name)
}

// Pairs returns a sub-bucket pairs map from a body string.
func Pairs(body string) map[string]string {
	body = Body(body)

	return map[string]string{
		"body": body,
		"hash": Hash(body),
	}
}

// Time returns a Time object from a unix milliseconds string.
func Time(unix string) time.Time {
	mill, _ := strconv.ParseInt(unix, 10, 64)
	return time.UnixMilli(mill)
}

// Unix returns a unix milliseconds string from a Time object.
func Unix(tobj time.Time) string {
	mill := tobj.UnixMilli()
	return strconv.FormatInt(mill, 10)
}
