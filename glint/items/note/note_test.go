package note

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/tools/bolt"
	"github.com/stvmln86/glint/glint/tools/neat"
	"github.com/stvmln86/glint/glint/tools/test"
)

func xNote(t *testing.T) *Note {
	db := test.MockDB(t)
	return &Note{db, "alpha"}
}

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB(t)
	init := neat.Unix(time.Now())

	// success
	note, err := Create(db, "charlie", "Charlie.\n")
	assert.NotNil(t, note.DB)
	assert.Equal(t, "charlie", note.Name)
	assert.NoError(t, err)

	// success - check database
	pairs, err := bolt.Get(db, "charlie", init)
	assert.Equal(t, "Charlie.\n", pairs["body"])
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	note, err := Get(db, "alpha")
	assert.NotNil(t, note.DB)
	assert.Equal(t, "alpha", note.Name)
	assert.NoError(t, err)

	// error - does not exist
	note, err = Get(db, "nope")
	assert.Nil(t, note)
	assert.EqualError(t, err, "cannot get note nope - does not exist")
}

func TestDelete(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	err := note.Delete()
	assert.NoError(t, err)

	// success - check database
	ok, err := bolt.Exists(note.DB, "alpha")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	ok, err := note.Exists()
	assert.True(t, ok)
	assert.NoError(t, err)
}
