// Package comms implements the Command interface and functions.
package comms

import (
	"fmt"
	"io"
	"strings"

	"github.com/stvmln86/glint/glint/items/book"
	"github.com/stvmln86/glint/glint/tools/clui"
)

// Command is a callable user-facing command.
type Command struct {
	Name  string
	Help  string
	Paras []string
	Func  CommandFunc
}

// CommandFunc is a function that peforms a Command's logic.
type CommandFunc func(io.Writer, *book.Book, map[string]string) error

// Commands is a slice of all existing Command types.
var Commands = map[string]*Command{
	"create": Create,
}

// New returns a new Command.
func New(name, paras, help string, cfun CommandFunc) *Command {
	return &Command{name, help, strings.Fields(paras), cfun}
}

// Run executes an existing Command with an argument slice.
func Run(w io.Writer, book *book.Book, elems []string) error {
	name, elems := clui.Split(elems)
	comm, ok := Commands[name]
	if !ok {
		return fmt.Errorf("cannot run command %q - does not exist", name)
	}

	pairs, err := clui.Parse(comm.Paras, elems)
	if err != nil {
		return err
	}

	return comm.Func(w, book, pairs)
}
