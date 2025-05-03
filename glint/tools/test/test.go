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
		"body": "Alpha.\n",
		"hash": "cd9fc009cdc95c830fd057df66c0c363d476f4dad534a2125fda9373d357e702",
		"init": "1000",
	},

	"bravo": {
		"body": "Bravo.\n",
		"hash": "8411871deeec869d84512c3414dca33589881ebac95a886bc4bd5a4e172bfe7c",
		"init": "2000",
	},
}

// MockDB returns a temporary database populated with MockData.
func MockDB(t *testing.T) *bbolt.DB {
	dest := filepath.Join(t.TempDir(), "mock.db")
	db, _ := bbolt.Open(dest, 0777, nil)

	db.Update(func(tx *bbolt.Tx) error {
		for name, pairs := range MockData {
			bckt, _ := tx.CreateBucket([]byte(name))

			for attr, data := range pairs {
				bckt.Put([]byte(attr), []byte(data))
			}
		}

		return nil
	})

	return db
}
