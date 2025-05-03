package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	assert.NotNil(t, db)

	// success - check database
	db.View(func(tx *bbolt.Tx) error {
		for name, pairs := range MockData {
			bckt := tx.Bucket([]byte(name))

			for attr, want := range pairs {
				data := bckt.Get([]byte(attr))
				assert.Equal(t, want, string(data))
			}
		}

		return nil
	})
}
