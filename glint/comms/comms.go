// Package comms implements the Command interface and functions.
package comms

import (
	"fmt"
	"io"

	"github.com/stvmln86/glint/glint/items/book"
	"github.com/stvmln86/glint/glint/tools/clui"
)

type Command interface {
	// Name returns the Command's callable name.
	Name() string

	// Help returns the Command's help string.
	Help() string

	// Paras returns the Command's argument parameters.
	Paras() []string

	// Run executes the Command's logic.
	Run(io.Writer, *book.Book, map[string]string) error
}

// CommandFunc is a function that returns an initialised Command.
type CommandFunc func() (Command, error)

// Commands is a map of all existing Command types.
var Commands = map[string]CommandFunc{
	// "create": create.New,
}

// Run executes an existing Command with an argument slice.
func Run(w io.Writer, book *book.Book, elems []string) error {
	name, elems := clui.Split(elems)
	cfun, ok := Commands[name]
	if !ok {
		return fmt.Errorf("cannot run command %q - does not exist", name)
	}

	comm, err := cfun()
	if err != nil {
		return err
	}

	pairs, err := clui.Parse(comm.Paras(), elems)
	if err != nil {
		return err
	}

	return comm.Run(w, book, pairs)
}
