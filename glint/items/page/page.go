// Package page implements the Page type and methods.
package page

import (
	"fmt"
	"time"

	"github.com/stvmln86/glint/glint/tools/bolt"
	"github.com/stvmln86/glint/glint/tools/neat"
	"go.etcd.io/bbolt"
)

// Page is a single historical note version in a database.
type Page struct {
	DB   *bbolt.DB
	Name string
	Body string
	Hash string
	Init string
}

// Create creates and returns a new Page.
func Create(db *bbolt.DB, name, body string) (*Page, error) {
	name = neat.Name(name)
	ok, err := bolt.Exists(db, name)

	switch {
	case !ok:
		return nil, fmt.Errorf("cannot append note %s - does not exist", name)
	case err != nil:
		return nil, fmt.Errorf("cannot append note %s - %w", name, err)
	}

	init := neat.Unix(time.Now())
	pairs := neat.Pairs(body)

	if err := bolt.Set(db, name, init, pairs); err != nil {
		return nil, fmt.Errorf("cannot create page %s/%s - %w", name, init, err)
	}

	return Get(db, name, init)
}

// Get returns an existing Page.
func Get(db *bbolt.DB, name, init string) (*Page, error) {
	name = neat.Name(name)
	ok, err := bolt.Exists(db, name, init)

	switch {
	case !ok:
		return nil, fmt.Errorf("cannot get page %s/%s - does not exist", name, init)
	case err != nil:
		return nil, fmt.Errorf("cannot get page %s/%s - %w", name, init, err)
	}

	pairs, err := bolt.Get(db, name, init)
	if err != nil {
		return nil, fmt.Errorf("cannot get page %s/%s - %w", name, init, err)
	}

	return &Page{db, name, pairs["body"], pairs["hash"], init}, nil
}

// Delete deletes the existing Page.
func (p *Page) Delete() error {
	if err := bolt.Delete(p.DB, p.Name, p.Init); err != nil {
		return fmt.Errorf("cannot delete page %s/%s - %w", p.Name, p.Init, err)
	}

	return nil
}

// Exists returns true if the Page exists.
func (p *Page) Exists() (bool, error) {
	ok, err := bolt.Exists(p.DB, p.Name, p.Init)
	if err != nil {
		return false, fmt.Errorf("cannot check page %s/%s - %w", p.Name, p.Init, err)
	}

	return ok, nil
}
