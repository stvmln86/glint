package book

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/items/note"
	"github.com/stvmln86/glint/glint/tools/test"
)

func mockBook(t *testing.T) *Book {
	dire := test.MockDire(t)
	return New(dire, ".extn", 0777)
}

func TestNew(t *testing.T) {
	// setup
	dire := test.MockDire(t)

	// success
	book := New(dire, ".extn", 0777)
	assert.Equal(t, dire, book.Dire)
	assert.Equal(t, ".extn", book.Extn)
	assert.Equal(t, os.FileMode(0777), book.Mode)
}

func TestCreate(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	note, err := book.Create("name", "Body.\n")
	assert.NoError(t, err)
	assert.Contains(t, note.Orig, "name.extn")
	test.AssertFile(t, note.Orig, "Body.\n")
}

func TestGet(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	note, err := book.Get("alpha")
	assert.Contains(t, note.Orig, "alpha.extn")
	assert.NoError(t, err)

	// error - does not exist
	note, err = book.Get("nope")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot get file "nope.extn" - does not exist`)
}

func TestGetOrCreate(t *testing.T) {
	// setup
	book := mockBook(t)

	// success - create
	note, err := book.GetOrCreate("name")
	assert.Contains(t, note.Orig, "name.extn")
	test.AssertFile(t, note.Orig, "\n")
	assert.NoError(t, err)

	// success - get
	note, err = book.GetOrCreate("name")
	assert.Contains(t, note.Orig, "name.extn")
	assert.NoError(t, err)
}

func TestFilter(t *testing.T) {
	// setup
	book := mockBook(t)
	ffun := func(note *note.Note) (bool, error) {
		return note.Name() == "alpha", nil
	}

	// success
	notes, err := book.Filter(ffun)
	assert.Len(t, notes, 1)
	assert.Contains(t, notes[0].Orig, "alpha.extn")
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes := book.List()
	assert.Len(t, notes, 2)
	assert.Contains(t, notes[0].Orig, "alpha.extn")
	assert.Contains(t, notes[1].Orig, "bravo.extn")
}

func TestMatch(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes := book.Match("ALPHA")
	assert.Len(t, notes, 1)
	assert.Contains(t, notes[0].Orig, "alpha.extn")
}

func TestSearch(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes, err := book.Search("ALPHA")
	assert.Len(t, notes, 1)
	assert.Contains(t, notes[0].Orig, "alpha.extn")
	assert.NoError(t, err)
}
