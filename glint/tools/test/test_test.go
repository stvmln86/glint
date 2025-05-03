package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestGet(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	data := Get(db, "alpha", "1000.000", "name")
	assert.Equal(t, "alpha", data)
}

func TestSet(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	Set(db, "alpha", "1000.000", "name", "test")
	data := Get(db, "alpha", "1000.000", "name")
	assert.Equal(t, "test", data)
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	assert.NotNil(t, db)

	// success - check mock data
	db.View(func(tx *bbolt.Tx) error {
		for buck, items := range MockData {
			bobj := tx.Bucket([]byte(buck))

			for subb, pairs := range items {
				sobj := bobj.Bucket([]byte(subb))

				for attr, want := range pairs {
					data := sobj.Get([]byte(attr))
					assert.Equal(t, want, string(data))
				}
			}
		}

		return nil
	})
}
