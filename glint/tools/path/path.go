// Package path implements file path manipulation functions.
package path

import (
	"path/filepath"
	"strings"
)

// Base returns a path's base name.
func Base(orig string) string {
	return filepath.Base(orig)
}

// Dire returns a path's parent directory.
func Dire(orig string) string {
	return filepath.Dir(orig)
}

// Extn returns a path's extension with a leading dot.
func Extn(orig string) string {
	base := Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[clip:]
	}

	return ""
}

// Join returns a joined path from a directory, name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Match returns true if a path's name contains a case-insensitive substring.
func Match(orig, subs string) bool {
	subs = strings.ToLower(subs)
	name := strings.ToLower(Name(orig))
	return strings.Contains(name, subs)
}

// Reextn returns a path with a different extension.
func Reextn(orig, extn string) string {
	dire := Dire(orig)
	name := Name(orig)
	return Join(dire, name, extn)
}

// Rename returns a path with a different name.
func Rename(orig, name string) string {
	dire := Dire(orig)
	extn := Extn(orig)
	return Join(dire, name, extn)
}

// Name returns a path's name.
func Name(orig string) string {
	base := Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[:clip]
	}

	return base
}
