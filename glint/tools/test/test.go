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
		"1000.000": {
			"name": "alpha",
			"body": "Alpha (old).",
			"hash": "2038ace1dd703f9661fd9657b7a842257c0082e640103d702f31f0264aa39050",
			"init": "1000.000",
		},

		"1100.000": {
			"name": "alpha",
			"body": "Alpha (new).",
			"hash": "c4eb3161e551ad770fe506df8e76cae3f36785a243482edae532703100a103ea",
			"init": "1100.000",
		},
	},

	"bravo": {
		"2000.000": {
			"name": "bravo",
			"body": "Bravo.",
			"hash": "65c1f9293df813a992d69f4fb83d430530e40dc8630409b30d0f8e58b07b1e14",
			"init": "2000.000",
		},
	},
}

// Get returns a sub-bucket value from a database.
func Get(dbse *bbolt.DB, buck, subb, attr string) string {
	var data string

	dbse.View(func(tx *bbolt.Tx) error {
		bobj := tx.Bucket([]byte(buck))
		sobj := bobj.Bucket([]byte(subb))
		data = string(sobj.Get([]byte(attr)))
		return nil
	})

	return data
}

// Set sets a sub-bucket value into a database.
func Set(dbse *bbolt.DB, buck, subb, attr, data string) {
	dbse.Update(func(tx *bbolt.Tx) error {
		bobj := tx.Bucket([]byte(buck))
		sobj := bobj.Bucket([]byte(subb))
		sobj.Put([]byte(attr), []byte(data))
		return nil
	})
}

// MockDB returns a temporary database populated with MockData.
func MockDB(t *testing.T) *bbolt.DB {
	dest := filepath.Join(t.TempDir(), "mock.db")
	dbse, _ := bbolt.Open(dest, 0777, nil)

	dbse.Update(func(tx *bbolt.Tx) error {
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

	return dbse
}
