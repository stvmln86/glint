// Package note implements the Note type and methods.
package note

import (
	"fmt"
	"time"

	"github.com/stvmln86/glint/glint/tools/bolt"
	"github.com/stvmln86/glint/glint/tools/neat"
	"go.etcd.io/bbolt"
)

// Note is a single versioned note entry in a database.
type Note struct {
	DB   *bbolt.DB
	Name string
}

// Create creates and returns a new Note.
func Create(db *bbolt.DB, name, body string) (*Note, error) {
	name = neat.Name(name)
	init := neat.Unix(time.Now())
	pairs := neat.Pairs(body)

	if err := bolt.Set(db, name, init, pairs); err != nil {
		return nil, fmt.Errorf("cannot create note %s - %w", name, err)
	}

	return Get(db, name)
}

// Get returns an existing Note.
func Get(db *bbolt.DB, name string) (*Note, error) {
	name = neat.Name(name)
	ok, err := bolt.Exists(db, name)

	switch {
	case !ok:
		return nil, fmt.Errorf("cannot get note %s - does not exist", name)
	case err != nil:
		return nil, fmt.Errorf("cannot get note %s - %w", name, err)
	}

	return &Note{db, name}, nil
}

// Delete deletes the existing Note.
func (n *Note) Delete() error {
	if err := bolt.Delete(n.DB, n.Name); err != nil {
		return fmt.Errorf("cannot delete note %s - %w", n.Name, err)
	}

	return nil
}

// Exists returns true if the Note exists.
func (n *Note) Exists() (bool, error) {
	ok, err := bolt.Exists(n.DB, n.Name)
	if err != nil {
		return false, fmt.Errorf("cannot check note %s - %w", n.Name, err)
	}

	return ok, nil
}
