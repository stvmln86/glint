// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
	"unicode"
)

// Body returns a whitespace-trimmed body string.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Extn returns a lowercase extension string with a leading dot.
func Extn(extn string) string {
	extn = strings.ToLower(extn)
	extn = strings.TrimSpace(extn)
	return "." + strings.TrimPrefix(extn, ".")
}

// Path returns a whitespace-trimmed clean file path.
func Path(path string) string {
	path = strings.TrimSpace(path)
	return filepath.Clean(path)
}

// Name returns a lowercase alphanumeric (with dashes) name string.
func Name(name string) string {
	var chars []rune
	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(char) || unicode.IsNumber(char):
			chars = append(chars, char)
		case unicode.IsSpace(char):
			chars = append(chars, '-')
		}
	}

	return strings.Trim(string(chars), "-")
}
