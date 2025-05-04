package book

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/items/note"
	"github.com/stvmln86/glint/glint/tools/bolt"
	"github.com/stvmln86/glint/glint/tools/neat"
	"github.com/stvmln86/glint/glint/tools/test"
)

func TestNew(t *testing.T) {
	// setup
	db := test.MockDB(t)
	path := db.Path()
	db.Close()

	// success
	book, err := New(path)
	assert.NotNil(t, book.DB)
	assert.NotEmpty(t, book.Path)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB(t)
	hash := neat.Hash("Charlie.\n")
	init := neat.Unix(time.Now())
	book := &Book{DB: db, Path: "test.db"}

	// success
	note, err := book.Create("charlie", "Charlie.\n")
	assert.Equal(t, "charlie", note.Name)
	assert.Equal(t, "Charlie.\n", note.Body)
	assert.Equal(t, hash, note.Hash)
	assert.Equal(t, init, note.Init)
	assert.NoError(t, err)

	// success - check database
	pairs, ok, err := bolt.Get(db, "charlie")
	assert.Equal(t, "Charlie.\n", pairs["body"])
	assert.Equal(t, hash, pairs["hash"])
	assert.Equal(t, init, pairs["init"])
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestFilter(t *testing.T) {
	// setup
	db := test.MockDB(t)
	book := &Book{DB: db, Path: "test.db"}
	ffun := func(note *note.Note) (bool, error) {
		return note.Name == "alpha", nil
	}

	// success
	notes, err := book.Filter(ffun)
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)
	book := &Book{DB: db, Path: "test.db"}

	// success
	note, err := book.Get("alpha")
	assert.NotNil(t, note.DB)
	assert.Equal(t, "alpha", note.Name)
	assert.Equal(t, "Alpha.\n", note.Body)
	assert.Equal(t, test.MockData["alpha"]["hash"], note.Hash)
	assert.Equal(t, "1000", note.Init)
	assert.NoError(t, err)
}

func TestGetOrCreate(t *testing.T) {
	// setup
	db := test.MockDB(t)
	hash := neat.Hash("Charlie.\n")
	init := neat.Unix(time.Now())
	book := &Book{DB: db, Path: "test.db"}

	// success - new note
	note, err := book.GetOrCreate("charlie", "Charlie.\n")
	assert.Equal(t, "charlie", note.Name)
	assert.Equal(t, "Charlie.\n", note.Body)
	assert.Equal(t, hash, note.Hash)
	assert.Equal(t, init, note.Init)
	assert.NoError(t, err)

	// success - check database
	pairs, ok, err := bolt.Get(db, "charlie")
	assert.Equal(t, "Charlie.\n", pairs["body"])
	assert.Equal(t, hash, pairs["hash"])
	assert.Equal(t, init, pairs["init"])
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - existing note
	note, err = book.GetOrCreate("charlie", "")
	assert.Equal(t, "charlie", note.Name)
	assert.Equal(t, "Charlie.\n", note.Body)
	assert.Equal(t, hash, note.Hash)
	assert.Equal(t, init, note.Init)
}

func TestList(t *testing.T) {
	// setup
	db := test.MockDB(t)
	book := &Book{DB: db, Path: "test.db"}

	// success
	notes, err := book.List()
	assert.Len(t, notes, 2)
	assert.Equal(t, "alpha", notes[0].Name)
	assert.Equal(t, "bravo", notes[1].Name)
	assert.NoError(t, err)
}
