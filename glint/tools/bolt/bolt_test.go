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

	// success - sub-bucket
	err = Delete(db, "bravo", "2000")
	assert.NoError(t, err)

	// success - check database
	db.View(func(tx *bbolt.Tx) error {
		bobj := tx.Bucket([]byte("bravo"))
		assert.NotNil(t, bobj)

		sobj := bobj.Bucket([]byte("2000"))
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

	// success - buckets
	bucks, err := List(db)
	assert.Equal(t, []string{"alpha", "bravo"}, bucks)
	assert.NoError(t, err)

	// success - sub-buckets
	subbs, err := List(db, "alpha")
	assert.Equal(t, []string{"1000", "1100"}, subbs)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	err := Set(db, "charlie", "3000", map[string]string{"body": "Charlie.\n"})
	assert.NoError(t, err)

	// success - check database
	db.View(func(tx *bbolt.Tx) error {
		bobj := tx.Bucket([]byte("charlie"))
		assert.NotNil(t, bobj)

		sobj := bobj.Bucket([]byte("3000"))
		assert.NotNil(t, sobj)

		body := sobj.Get([]byte("body"))
		assert.Equal(t, "data", string(body))
		return nil
	})
}
