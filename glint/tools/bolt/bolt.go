// Package bolt implements database handling functions.
package bolt

import "go.etcd.io/bbolt"

// Delete deletes an existing bucket from a database.
func Delete(db *bbolt.DB, name string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(name))
	})
}

// Exists returns true if a bucket exists in a database.
func Exists(db *bbolt.DB, name string) (bool, error) {
	var ok bool

	return ok, db.View(func(tx *bbolt.Tx) error {
		ok = tx.Bucket([]byte(name)) != nil
		return nil
	})
}

// Get returns an existing bucket from a database.
func Get(db *bbolt.DB, name string) (map[string]string, error) {
	var pairs = make(map[string]string)

	return pairs, db.View(func(tx *bbolt.Tx) error {
		if bckt := tx.Bucket([]byte(name)); bckt != nil {
			return bckt.ForEach(func(attr, data []byte) error {
				pairs[string(attr)] = string(data)
				return nil
			})
		}

		return nil
	})
}

// List returns all existing bucket names in a database.
func List(db *bbolt.DB) ([]string, error) {
	var names []string

	return names, db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			names = append(names, string(name))
			return nil
		})
	})
}

// Set sets a new or existing bucket into a database.
func Set(db *bbolt.DB, name string, pairs map[string]string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bckt, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		for attr, data := range pairs {
			if err := bckt.Put([]byte(attr), []byte(data)); err != nil {
				return err
			}
		}

		return nil
	})
}
