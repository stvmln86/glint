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

// Join returns a joined path from a directory, slug and extension.
func Join(dire, slug, extn string) string {
	return filepath.Join(dire, slug+extn)
}

// Match returns true if a path's slug contains a case-insensitive substring.
func Match(orig, subs string) bool {
	subs = strings.ToLower(subs)
	slug := strings.ToLower(Slug(orig))
	return strings.Contains(slug, subs)
}

// Reextn returns a path with a changed extension.
func Reextn(orig, extn string) string {
	dire := Dire(orig)
	slug := Slug(orig)
	return Join(dire, slug, extn)
}

// Reslug returns a path with a changed slug.
func Reslug(orig, slug string) string {
	dire := Dire(orig)
	extn := Extn(orig)
	return Join(dire, slug, extn)
}

// Slug returns a path's slug.
func Slug(orig string) string {
	base := Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[:clip]
	}

	return ""
}
