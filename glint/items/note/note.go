// Package note implements the Note type and methods.
package note

import (
	"os"

	"github.com/stvmln86/glint/glint/tools/file"
	"github.com/stvmln86/glint/glint/tools/neat"
	"github.com/stvmln86/glint/glint/tools/path"
)

// Note is a single plaintext note file in a Book.
type Note struct {
	Orig string
	Mode os.FileMode
}

// New returns a new Note.
func New(orig string, mode os.FileMode) *Note {
	return &Note{orig, mode}
}

// Delete "deletes" the Note by changing its extension to ".trash".
func (n *Note) Delete() error {
	return file.Delete(n.Orig)
}

// Exists returns true if the Note exists.
func (n *Note) Exists() bool {
	return file.Exists(n.Orig)
}

// Match returns true if the Note's name contains a case-insensitive substring.
func (n *Note) Match(subs string) bool {
	return path.Match(n.Orig, subs)
}

// Name returns the Note's name.
func (n *Note) Name() string {
	name := path.Name(n.Orig)
	return neat.Name(name)
}

// Read returns the Note's body as a string.
func (n *Note) Read() (string, error) {
	body, err := file.Read(n.Orig)
	return neat.Body(body), err
}

// Rename renames the Note to a different name.
func (n *Note) Rename(name string) error {
	name = neat.Name(name)
	return file.Rename(n.Orig, name)
}

// Search returns true if the Note's body contains a case-insensitive substring.
func (n *Note) Search(subs string) (bool, error) {
	return file.Search(n.Orig, subs)
}

// Update overwrites the Note's body with a string.
func (n *Note) Update(body string) error {
	body = neat.Body(body)
	return file.Update(n.Orig, body, n.Mode)
}
