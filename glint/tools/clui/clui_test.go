package clui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	// setup
	os.Setenv("TEST", "\tTest.\n")
	os.Setenv("BLANK", "\n")

	// success
	data, err := Env("TEST")
	assert.Equal(t, "Test.", data)
	assert.NoError(t, err)

	// error - does not exist
	data, err = Env("NOPE")
	assert.Empty(t, data)
	assert.EqualError(t, err, `cannot find envvar "NOPE" - does not exist`)

	// error - is blank
	data, err = Env("BLANK")
	assert.Empty(t, data)
	assert.EqualError(t, err, `cannot find envvar "BLANK" - is blank`)
}

func TestParse(t *testing.T) {
	// success - with argument
	pairs, err := Parse([]string{"parameter:default"}, []string{"argument"})
	assert.Equal(t, map[string]string{"parameter": "argument"}, pairs)
	assert.NoError(t, err)

	// success - with default
	pairs, err = Parse([]string{"parameter:default"}, nil)
	assert.Equal(t, map[string]string{"parameter": "default"}, pairs)
	assert.NoError(t, err)

	// error - cannot parse arguments
	pairs, err = Parse([]string{"parameter"}, nil)
	assert.Empty(t, pairs)
	assert.EqualError(t, err, `cannot parse arguments - "parameter" not provided`)
}

func TestSplit(t *testing.T) {
	// success - zero arguments
	name, elems := Split(nil)
	assert.Empty(t, name)
	assert.Empty(t, elems)

	// success - one argument
	name, elems = Split([]string{"NAME"})
	assert.Equal(t, "name", name)
	assert.Empty(t, elems)

	// success - multiple arguments
	name, elems = Split([]string{"NAME", "argument"})
	assert.Equal(t, "name", name)
	assert.Equal(t, []string{"argument"}, elems)
}
