package note

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/tools/bolt"
	"github.com/stvmln86/glint/glint/tools/neat"
	"github.com/stvmln86/glint/glint/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB(t)
	hash := neat.Hash("Charlie.\n")
	init := neat.Unix(time.Now())

	// success
	note, err := Create(db, "charlie", "Charlie.\n")
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

	// error - already exists
	note, err = Create(db, "charlie", "Charlie.\n")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot create note "charlie" - already exists`)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	note, err := Get(db, "alpha")
	assert.NotNil(t, note.DB)
	assert.Equal(t, "alpha", note.Name)
	assert.Equal(t, "Alpha.\n", note.Body)
	assert.Equal(t, test.MockData["alpha"]["hash"], note.Hash)
	assert.Equal(t, "1000", note.Init)
	assert.NoError(t, err)

	// error - does not exist
	note, err = Get(db, "nope")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot get note "nope" - does not exist`)
}

func TestDelete(t *testing.T) {
	// setup
	db := test.MockDB(t)
	note, _ := Get(db, "alpha")

	// success
	err := note.Delete()
	assert.NoError(t, err)

	// success - check database
	ok, err := bolt.Exists(db, "alpha")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	db := test.MockDB(t)
	note, _ := Get(db, "alpha")
	fail := &Note{DB: db, Name: "nope"}

	// success - true
	ok, err := note.Exists()
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - false
	ok, err = fail.Exists()
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	db := test.MockDB(t)
	note, _ := Get(db, "alpha")
	hash := neat.Hash("Body.\n")

	// success
	err := note.Update("Body.\n")
	assert.Equal(t, "alpha", note.Name)
	assert.Equal(t, "Body.\n", note.Body)
	assert.Equal(t, hash, note.Hash)
	assert.Equal(t, "1000", note.Init)
	assert.NoError(t, err)

	// success - check database
	pairs, ok, err := bolt.Get(db, "alpha")
	assert.Equal(t, "Body.\n", pairs["body"])
	assert.Equal(t, hash, pairs["hash"])
	assert.Equal(t, "1000", pairs["init"])
	assert.True(t, ok)
	assert.NoError(t, err)
}
