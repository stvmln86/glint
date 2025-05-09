package comms

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/items/book"
	"github.com/stvmln86/glint/glint/tools/test"
)

var MockCommand = New(
	"mock", "parameter:default",
	"A mock command for unit testing.",

	func(w io.Writer, _ *book.Book, pairs map[string]string) error {
		fmt.Fprintf(w, "parameter=%s\n", pairs["parameter"])
		return nil
	},
)

func runCommand(t *testing.T, elems ...string) (*bytes.Buffer, *book.Book, error) {
	w := new(bytes.Buffer)
	dire := test.MockDire(t)
	book := book.New(dire, ".extn", 0666)
	return w, book, Run(w, book, elems)
}

func TestRun(t *testing.T) {
	// setup
	w := bytes.NewBuffer(nil)
	Commands["mock"] = MockCommand

	// success
	err := Run(w, nil, []string{"mock", "argument"})
	assert.NoError(t, err)
	assert.Equal(t, "parameter=argument\n", w.String())

	// error - does not exist
	err = Run(nil, nil, []string{"nope"})
	assert.EqualError(t, err, `cannot run command "nope" - does not exist`)
}
