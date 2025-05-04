// Package book implements the Book type and methods.
package book

import (
	"fmt"

	"github.com/stvmln86/glint/glint/items/note"
	"github.com/stvmln86/glint/glint/tools/bolt"
	"github.com/stvmln86/glint/glint/tools/neat"
	"go.etcd.io/bbolt"
)

// Book is a database of plaintext Notes.
type Book struct {
	DB   *bbolt.DB
	Path string
}

// New returns a new Book.
func New(path string) (*Book, error) {
	db, err := bbolt.Open(path, 0777, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot open book %q - %w", path, err)
	}

	return &Book{db, path}, nil
}

// Create creates and returns a new Note in the Book.
func (b *Book) Create(name, body string) (*note.Note, error) {
	return note.Create(b.DB, name, body)
}

// Filter returns all Notes in the Book that pass a filter function.
func (b *Book) Filter(ffun func(*note.Note) (bool, error)) ([]*note.Note, error) {
	var goods []*note.Note

	notes, err := b.List()
	if err != nil {
		return nil, err
	}

	for _, note := range notes {
		ok, err := ffun(note)
		switch {
		case err != nil:
			return nil, err
		case ok:
			goods = append(goods, note)
		}
	}

	return goods, nil
}

// Get returns an existing Note from the Book.
func (b *Book) Get(name string) (*note.Note, error) {
	return note.Get(b.DB, name)
}

// GetOrCreate returns a newly created or existing Note from the Book.
func (b *Book) GetOrCreate(name, body string) (*note.Note, error) {
	name = neat.Name(name)
	ok, err := bolt.Exists(b.DB, name)

	switch {
	case !ok:
		return note.Create(b.DB, name, body)
	case err != nil:
		return nil, err
	default:
		return note.Get(b.DB, name)
	}
}

// List returns all Notes in the Book.
func (b *Book) List() ([]*note.Note, error) {
	names, err := bolt.List(b.DB)
	if err != nil {
		return nil, fmt.Errorf("cannot list book %q - %w", b.Path, err)
	}

	var notes []*note.Note
	for _, name := range names {
		note, err := note.Get(b.DB, name)
		if err != nil {
			return nil, fmt.Errorf("cannot list book %q - %w", b.Path, err)
		}

		notes = append(notes, note)
	}

	return notes, nil
}
