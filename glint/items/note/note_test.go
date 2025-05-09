package note

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/tools/path"
	"github.com/stvmln86/glint/glint/tools/test"
)

func mockNote(t *testing.T) *Note {
	orig := test.MockFile(t, "alpha.extn")
	return New(orig, 0666)
}

func TestNew(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	note := New(orig, 0666)
	assert.Equal(t, note.Orig, orig)
	assert.Equal(t, os.FileMode(0666), note.Mode)
}

func TestDelete(t *testing.T) {
	// setup
	note := mockNote(t)
	dest := path.Reextn(note.Orig, ".trash")

	// success
	err := note.Delete()
	assert.NoFileExists(t, note.Orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok := note.Exists()
	assert.True(t, ok)
}

func TestMatch(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok := note.Match("ALPHA")
	assert.True(t, ok)
}

func TestName(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func TestRead(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, "Alpha.\n", body)
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	note := mockNote(t)
	dest := path.Rename(note.Orig, "name")

	// success
	err := note.Rename("name")
	assert.NoFileExists(t, note.Orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok, err := note.Search("ALPHA")
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	err := note.Update("Body.\n")
	test.AssertFile(t, note.Orig, "Body.\n")
	assert.NoError(t, err)
}
