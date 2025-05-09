// Package comms implements the Command interface and functions.
package comms

import (
	"fmt"
	"io"
	"strings"

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

// Get returns a CommandFunc from a disambiguated Command name.
func Get(part string) (CommandFunc, error) {
	var cfuns []CommandFunc
	for name, cfun := range Commands {
		if strings.HasPrefix(name, part) {
			cfuns = append(cfuns, cfun)
		}
	}

	switch len(cfuns) {
	case 0:
		return nil, fmt.Errorf("cannot run command %q - does not exist", part)
	case 1:
		return cfuns[0], nil
	default:
		return nil, fmt.Errorf("cannot run command %q - is ambiguous", part)
	}
}

// Run executes an existing Command with an argument slice.
func Run(w io.Writer, book *book.Book, elems []string) error {
	name, elems := clui.Split(elems)
	cfun, err := Get(name)
	if err != nil {
		return err
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
