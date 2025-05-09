package comms

import (
	"fmt"
	"io"

	"github.com/stvmln86/glint/glint/items/book"
)

// Create is a Command that creates Notes.
var Create = New(
	"create", "name body:",
	"Create a new note with an optional body.",

	func(w io.Writer, book *book.Book, pairs map[string]string) error {
		note, err := book.Create(pairs["name"], pairs["body"])
		if err != nil {
			return err
		}

		fmt.Fprintf(w, "Created note %q.\n", note.Name())
		return nil
	},
)
