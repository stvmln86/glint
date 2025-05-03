// Package note implements the Note type and methods.
package note

import (
	"fmt"

	"github.com/stvmln86/glint/glint/tools/bolt"
	"github.com/stvmln86/glint/glint/tools/neat"
	"go.etcd.io/bbolt"
)

// Note is a single plaintext note entry in a Book.
type Note struct {
	DB   *bbolt.DB
	Name string
	Body string
	Hash string
	Init string
}

// Get returns an existing Note by name.
func Get(db *bbolt.DB, name string) (*Note, error) {
	name = neat.Name(name)
	pairs, ok, err := bolt.Get(db, name)

	switch {
	case !ok:
		return nil, fmt.Errorf("cannot get note %q - does not exist", name)
	case err != nil:
		return nil, fmt.Errorf("cannot get note %q - %w", name, err)
	}

	return &Note{db, name, pairs["body"], pairs["hash"], pairs["init"]}, nil
}

// Delete deletes the existing Note.
func (n *Note) Delete() error {
	if err := bolt.Delete(n.DB, n.Name); err != nil {
		return fmt.Errorf("cannot delete note %q - %w", n.Name, err)
	}

	return nil
}

// Exists returns true if the Note exists.
func (n *Note) Exists() (bool, error) {
	ok, err := bolt.Exists(n.DB, n.Name)
	if err != nil {
		return false, fmt.Errorf("cannot check note %q - %w", n.Name, err)
	}

	return ok, nil
}

// Update updates the existing Note's body and hash.
func (n *Note) Update(body string) error {
	body = neat.Body(body)
	hash := neat.Hash(body)
	pairs := map[string]string{"body": body, "hash": hash}

	if err := bolt.Set(n.DB, n.Name, pairs); err != nil {
		return fmt.Errorf("cannot update note %q - %w", n.Name, err)
	}

	n.Body, n.Hash = body, hash
	return nil
}
