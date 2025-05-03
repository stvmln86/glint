// Package test implements unit testing helper functions and mock data.
package test

import (
	"path/filepath"
	"testing"

	"go.etcd.io/bbolt"
)

// MockData is a map of mock database data for unit testing.
var MockData = map[string]map[string]map[string]string{
	"alpha": {
		"1000": {
			"body": "Alpha (old).\n",
			"hash": "2038ace1dd703f9661fd9657b7a842257c0082e640103d702f31f0264aa39050",
		},

		"1100": {
			"body": "Alpha (new).\n",
			"hash": "c4eb3161e551ad770fe506df8e76cae3f36785a243482edae532703100a103ea",
		},
	},

	"bravo": {
		"2000": {
			"body": "Bravo.\n",
			"hash": "65c1f9293df813a992d69f4fb83d430530e40dc8630409b30d0f8e58b07b1e14",
		},
	},
}

// MockDB returns a temporary database populated with MockData.
func MockDB(t *testing.T) *bbolt.DB {
	dest := filepath.Join(t.TempDir(), "mock.db")
	db, _ := bbolt.Open(dest, 0777, nil)

	db.Update(func(tx *bbolt.Tx) error {
		for buck, items := range MockData {
			bobj, _ := tx.CreateBucket([]byte(buck))

			for subb, pairs := range items {
				sobj, _ := bobj.CreateBucket([]byte(subb))

				for attr, data := range pairs {
					sobj.Put([]byte(attr), []byte(data))
				}
			}
		}

		return nil
	})

	return db
}
