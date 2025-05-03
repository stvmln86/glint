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

	// success - check bucket deleted
	db.View(func(tx *bbolt.Tx) error {
		bobj := tx.Bucket([]byte("alpha"))
		assert.Nil(t, bobj)
		return nil
	})

	// success - sub-bucket
	err = Delete(db, "bravo", "3000")
	assert.NoError(t, err)

	// success - check sub-bucket deleted
	db.View(func(tx *bbolt.Tx) error {
		bobj := tx.Bucket([]byte("bravo"))
		assert.NotNil(t, bobj)
		sobj := bobj.Bucket([]byte("3000"))
		assert.Nil(t, sobj)
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

	// success - sub-bucket, true
	ok, err = Exists(db, "alpha", "1000")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - sub-bucket, false
	ok, err = Exists(db, "alpha", "nope")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	pairs, err := Get(db, "alpha", "1000")
	assert.Equal(t, test.MockData["alpha"]["1000"], pairs)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	subbs, err := List(db, "alpha")
	assert.Equal(t, []string{"1000", "2000"}, subbs)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	err := Set(db, "buck", "subb", map[string]string{"attr": "data"})
	assert.NoError(t, err)

	// success - check data set
	data := test.Get(db, "buck", "subb", "attr")
	assert.Equal(t, "data", data)
}
