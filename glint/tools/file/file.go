// Package file implements file system input/output functions.
package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/stvmln86/glint/glint/tools/path"
)

// Create creates a new file containing a string.
func Create(dest, body string, mode os.FileMode) error {
	base := path.Base(dest)

	if Exists(dest) {
		return fmt.Errorf("cannot create file %q - already exists", base)
	}

	if err := os.WriteFile(dest, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot create file %q - %w", base, err)
	}

	return nil
}

// Delete "deletes" an existing file by changing its extension to ".trash".
func Delete(orig string) error {
	base := path.Base(orig)

	if !Exists(orig) {
		return fmt.Errorf("cannot delete file %q - does not exist", base)
	}

	dest := path.Reextn(orig, ".trash")
	if err := os.Rename(orig, dest); err != nil {
		return fmt.Errorf("cannot delete file %q - %w", base, err)
	}

	return nil
}

// Exists returns true if a file exists.
func Exists(orig string) bool {
	_, err := os.Stat(orig)
	return !errors.Is(err, os.ErrNotExist)
}

// Glob returns a sorted slice of all paths in a directory matching an extension.
func Glob(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	origs, _ := filepath.Glob(glob)
	return origs
}

// Read returns an existing file's body as a string.
func Read(orig string) (string, error) {
	base := path.Base(orig)

	if !Exists(orig) {
		return "", fmt.Errorf("cannot read file %q - does not exist", base)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		return "", fmt.Errorf("cannot read file %q - %w", base, err)
	}

	return string(bytes), nil
}

// Rename renames an existing file to a different name.
func Rename(orig, name string) error {
	base := path.Base(orig)
	dest := path.Rename(orig, name)

	if !Exists(orig) {
		return fmt.Errorf("cannot rename file %q - does not exist", base)
	}

	if err := os.Rename(orig, dest); err != nil {
		return fmt.Errorf("cannot rename file %q - %w", base, err)
	}

	return nil
}

// Search returns true if an existing file's body contains a case-insensitive substring.
func Search(orig, subs string) (bool, error) {
	base := path.Base(orig)

	if !Exists(orig) {
		return false, fmt.Errorf("cannot search file %q - does not exist", base)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		return false, fmt.Errorf("cannot search file %q - %w", base, err)
	}

	subs = strings.ToLower(subs)
	body := strings.ToLower(string(bytes))
	return strings.Contains(body, subs), nil
}

// Update overwrites an existing file's body with a string.
func Update(orig, body string, mode os.FileMode) error {
	base := path.Base(orig)

	if !Exists(orig) {
		return fmt.Errorf("cannot update file %q - does not exist", base)
	}

	if err := os.WriteFile(orig, []byte(body), 0777); err != nil {
		return fmt.Errorf("cannot update file %q - %w", base, err)
	}

	return nil
}
