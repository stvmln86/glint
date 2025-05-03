package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/tools/test"
	"go.etcd.io/bbolt"
)

func TestDelete(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	err := Delete(db, "alpha")
	assert.NoError(t, err)

	// success - check database
	db.View(func(tx *bbolt.Tx) error {
		bckt := tx.Bucket([]byte("alpha"))
		assert.Nil(t, bckt)
		return nil
	})
}

func TestExists(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success - true
	ok, err := Exists(db, "alpha")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - false
	ok, err = Exists(db, "nope")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	pairs, err := Get(db, "alpha")
	assert.Equal(t, test.MockData["alpha"], pairs)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	names, err := List(db)
	assert.Equal(t, []string{"alpha", "bravo"}, names)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	err := Set(db, "charlie", map[string]string{"body": "Charlie.\n"})
	assert.NoError(t, err)

	// success - check database
	db.View(func(tx *bbolt.Tx) error {
		bckt := tx.Bucket([]byte("charlie"))
		data := bckt.Get([]byte("body"))
		assert.Equal(t, "Charlie.\n", string(data))
		return nil
	})
}
