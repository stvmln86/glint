// Package test implements unit testing helper functions and mock data.
package test

import (
	"path/filepath"
	"testing"

	"go.etcd.io/bbolt"
)

// MockData is a map of mock database data for unit testing.
var MockData = map[string]map[string]string{
	"alpha": {
		"1000": "Alpha one.\n",
		"1100": "Alpha two.\n",
	},

	"bravo": {
		"2000": "Bravo.\n",
	},
}

// MockDB returns a temporary database populated with MockData.
func MockDB(t *testing.T) *bbolt.DB {
	dest := filepath.Join(t.TempDir(), "mock.db")
	db, _ := bbolt.Open(dest, 0777, nil)

	db.Update(func(tx *bbolt.Tx) error {
		for buck, pairs := range MockData {
			bobj, _ := tx.CreateBucket([]byte(buck))

			for attr, data := range pairs {
				bobj.Put([]byte(attr), []byte(data))
			}
		}

		return nil
	})

	return db
}
