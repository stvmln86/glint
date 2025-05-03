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

	// success - bucket
	err := Delete(db, "alpha")
	assert.NoError(t, err)

	// success - check database
	db.View(func(tx *bbolt.Tx) error {
		bobj := tx.Bucket([]byte("alpha"))
		assert.Nil(t, bobj)
		return nil
	})

	// success - pair
	err = Delete(db, "bravo", "2000")
	assert.NoError(t, err)

	// success - check database
	db.View(func(tx *bbolt.Tx) error {
		bobj := tx.Bucket([]byte("bravo"))
		data := bobj.Get([]byte("2000"))
		assert.Nil(t, data)
		return nil
	})
}

func TestExists(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success - bucket, true
	ok, err := Exists(db, "alpha")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - bucket, false
	ok, err = Exists(db, "nope")
	assert.False(t, ok)
	assert.NoError(t, err)

	// success - pair, true
	ok, err = Exists(db, "alpha", "1000")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - pair, false
	ok, err = Exists(db, "alpha", "nope")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	pairs, err := Get(db, "alpha", "1000")
	assert.Equal(t, "Alpha one.\n", pairs)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success - buckets
	bucks, err := List(db)
	assert.Equal(t, []string{"alpha", "bravo"}, bucks)
	assert.NoError(t, err)

	// success - pairs
	subbs, err := List(db, "alpha")
	assert.Equal(t, []string{"1000", "1100"}, subbs)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	err := Set(db, "charlie", "3000", "Charlie.\n")
	assert.NoError(t, err)

	// success - check database
	db.View(func(tx *bbolt.Tx) error {
		bobj := tx.Bucket([]byte("charlie"))
		data := bobj.Get([]byte("3000"))
		assert.Equal(t, "Charlie.\n", string(data))
		return nil
	})
}
