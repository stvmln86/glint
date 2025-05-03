package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestGet(t *testing.T) {
	// setup
	dbse := MockDB(t)

	// success
	data := Get(dbse, "alpha", "1000.000", "name")
	assert.Equal(t, "alpha", data)
}

func TestSet(t *testing.T) {
	// setup
	dbse := MockDB(t)

	// success
	Set(dbse, "alpha", "1000.000", "name", "test")
	data := Get(dbse, "alpha", "1000.000", "name")
	assert.Equal(t, "test", data)
}

func TestMockDB(t *testing.T) {
	// success
	dbse := MockDB(t)
	assert.NotNil(t, dbse)

	// success - check data
	dbse.View(func(tx *bbolt.Tx) error {
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
