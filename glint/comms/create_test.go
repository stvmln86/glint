package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// success - no body
	w, book, err := runCommand(t, "create", "name")
	note, _ := book.Get("name")
	assert.NoError(t, err)
	assert.Equal(t, `Created note "name".`+"\n", w.String())
	assert.True(t, note.Exists())
}
