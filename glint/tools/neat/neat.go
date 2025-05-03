// Package neat implements data sanitisation and conversion functions.
package neat

import (
	"strconv"
	"strings"
	"time"
)

// Body returns a whitespace-trimmed body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Name returns a whitespace-trimmed lowercase name string.
func Name(name string) string {
	name = strings.ToLower(name)
	return strings.TrimSpace(name)
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
