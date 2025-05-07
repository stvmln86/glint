// Package book implements the Book type and methods.
package book

import (
	"fmt"
	"os"

	"github.com/stvmln86/glint/glint/items/note"
	"github.com/stvmln86/glint/glint/tools/file"
	"github.com/stvmln86/glint/glint/tools/neat"
	"github.com/stvmln86/glint/glint/tools/path"
)

// Book is a directory of plaintext Note files.
type Book struct {
	Dire string
	Extn string
	Mode os.FileMode
}

// New returns a new Book.
func New(dire, extn string, mode os.FileMode) *Book {
	dire = neat.Path(dire)
	extn = neat.Extn(extn)
	return &Book{dire, extn, mode}
}

// Create returns a newly created Note containing a string.
func (b *Book) Create(name, body string) (*note.Note, error) {
	name = neat.Name(name)
	body = neat.Body(body)
	dest := path.Join(b.Dire, name, b.Extn)

	if err := file.Create(dest, body, b.Mode); err != nil {
		return nil, err
	}

	return note.New(dest, b.Mode), nil
}

// Get returns an existing Note.
func (b *Book) Get(name string) (*note.Note, error) {
	name = neat.Name(name)
	orig := path.Join(b.Dire, name, b.Extn)
	base := path.Base(orig)

	if !file.Exists(orig) {
		return nil, fmt.Errorf("cannot get file %q - does not exist", base)
	}

	return note.New(orig, b.Mode), nil
}

// GetOrCreate returns a new (blank) or existing Note.
func (b *Book) GetOrCreate(name string) (*note.Note, error) {
	name = neat.Name(name)
	orig := path.Join(b.Dire, name, b.Extn)

	if !file.Exists(orig) {
		return b.Create(name, "\n")
	}

	return note.New(orig, b.Mode), nil
}

// Filter returns all Notes in the Book passing a filter function.
func (b *Book) Filter(ffun func(*note.Note) (bool, error)) ([]*note.Note, error) {
	var notes []*note.Note
	for _, note := range b.List() {
		ok, err := ffun(note)
		switch {
		case err != nil:
			return nil, err
		case ok:
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// List returns all existing Notes in the Book.
func (b *Book) List() []*note.Note {
	var notes []*note.Note

	for _, orig := range file.Glob(b.Dire, b.Extn) {
		note := note.New(orig, b.Mode)
		notes = append(notes, note)
	}

	return notes
}

// Match returns all Notes with names containing a case-insensitive substring.
func (b *Book) Match(subs string) []*note.Note {
	notes, _ := b.Filter(func(note *note.Note) (bool, error) {
		return note.Match(subs), nil
	})

	return notes
}

// Search returns all Notes with bodies containing a case-insensitive substring.
func (b *Book) Search(subs string) ([]*note.Note, error) {
	return b.Filter(func(note *note.Note) (bool, error) {
		return note.Search(subs)
	})
}
