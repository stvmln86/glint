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

func TestGet(t *testing.T) {
	// setup
	Commands = map[string]CommandFunc{
		"mock1": NewMockCommand,
		"mock2": NewMockCommand,
	}

	// success
	cfun, err := Get("mock1")
	assert.NotNil(t, cfun)
	assert.NoError(t, err)

	// error - does not exist
	cfun, err = Get("nope")
	assert.Nil(t, cfun)
	assert.EqualError(t, err, `cannot run command "nope" - does not exist`)

	// error - is ambiguous
	cfun, err = Get("mock")
	assert.Nil(t, cfun)
	assert.EqualError(t, err, `cannot run command "mock" - is ambiguous`)
}

func TestRun(t *testing.T) {
	// setup
	w := bytes.NewBuffer(nil)
	Commands = map[string]CommandFunc{
		"mock": NewMockCommand,
	}

	// success
	err := Run(w, nil, []string{"mock", "argument"})
	assert.NoError(t, err)
	assert.Equal(t, "parameter=argument\n", w.String())

	// error - does not exist
	err = Run(nil, nil, []string{"nope"})
	assert.EqualError(t, err, `cannot run command "nope" - does not exist`)
}
