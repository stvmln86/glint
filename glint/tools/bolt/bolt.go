// Package bolt implements database handling functions.
package bolt

import "go.etcd.io/bbolt"

// Delete deletes an existing bucket or pair  from a database.
func Delete(db *bbolt.DB, buck string, attr ...string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		if len(attr) == 0 {
			return tx.DeleteBucket([]byte(buck))
		}

		if bobj := tx.Bucket([]byte(buck)); bobj != nil {
			return bobj.Delete([]byte(attr[0]))
		}

		return nil
	})
}

// Exists returns true if a bucket or pair exists in a database.
func Exists(db *bbolt.DB, buck string, attr ...string) (bool, error) {
	var ok bool

	return ok, db.View(func(tx *bbolt.Tx) error {
		if len(attr) == 0 {
			ok = tx.Bucket([]byte(buck)) != nil
			return nil
		}

		if bobj := tx.Bucket([]byte(buck)); bobj != nil {
			ok = bobj.Get([]byte(attr[0])) != nil
		}

		return nil
	})
}

// Get returns an existing bucket pair's value from from a database.
func Get(db *bbolt.DB, buck, attr string) (string, error) {
	var data string

	return data, db.View(func(tx *bbolt.Tx) error {
		if bobj := tx.Bucket([]byte(buck)); bobj != nil {
			data = string(bobj.Get([]byte(attr)))
		}

		return nil
	})
}

// List returns all existing bucket names or bucket pair names from a database.
func List(db *bbolt.DB, buck ...string) ([]string, error) {
	var names []string

	return names, db.View(func(tx *bbolt.Tx) error {
		if len(buck) == 0 {
			return tx.ForEach(func(buck []byte, _ *bbolt.Bucket) error {
				names = append(names, string(buck))
				return nil
			})
		}

		if bobj := tx.Bucket([]byte(buck[0])); bobj != nil {
			return bobj.ForEach(func(attr, _ []byte) error {
				names = append(names, string(attr))
				return nil
			})
		}

		return nil
	})
}

// Set sets a pair in a new or existing bucket in a database.
func Set(db *bbolt.DB, buck, attr, data string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bobj, err := tx.CreateBucketIfNotExists([]byte(buck))
		if err != nil {
			return err
		}

		return bobj.Put([]byte(attr), []byte(data))
	})
}
