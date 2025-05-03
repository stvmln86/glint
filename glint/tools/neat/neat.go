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
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// Name returns a whitespace-trimmed lowercase name string.
func Name(name string) string {
	name = strings.ToLower(name)
	return strings.TrimSpace(name)
}

// Time returns a Time object from a unix seconds string.
func Time(unix string) time.Time {
	secs, _ := strconv.ParseInt(unix, 10, 64)
	return time.Unix(secs, 0)
}

// Unix returns a unix seconds string from a Time object.
func Unix(tobj time.Time) string {
	secs := tobj.Unix()
	return strconv.FormatInt(secs, 10)
}
