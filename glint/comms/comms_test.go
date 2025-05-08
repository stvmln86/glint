package comms

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/items/book"
)

type MockCommand struct{}

func NewMockCommand() (Command, error) {
	return &MockCommand{}, nil
}

func (c *MockCommand) Name() string {
	return "mock"
}

func (c *MockCommand) Help() string {
	return "A mock command for unit testing."
}

func (c *MockCommand) Paras() []string {
	return []string{"parameter:default"}
}

func (c *MockCommand) Run(w io.Writer, _ *book.Book, pairs map[string]string) error {
	fmt.Fprintf(w, "parameter=%s\n", pairs["parameter"])
	return nil
}

func TestRun(t *testing.T) {
	// setup
	w := bytes.NewBuffer(nil)
	Commands["mock"] = NewMockCommand

	// success
	err := Run(w, nil, []string{"mock", "argument"})
	assert.NoError(t, err)
	assert.Equal(t, "parameter=argument\n", w.String())

	// error - does not exist
	err = Run(nil, nil, []string{"nope"})
	assert.EqualError(t, err, `cannot run command "nope" - does not exist`)
}
