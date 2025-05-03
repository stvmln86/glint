// Package bolt implements database handling functions.
package bolt

import "go.etcd.io/bbolt"

// Delete deletes an existing bucket or sub-bucket from a database.
func Delete(db *bbolt.DB, buck string, subb ...string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		if len(subb) == 0 {
			return tx.DeleteBucket([]byte(buck))
		}

		if bobj := tx.Bucket([]byte(buck)); bobj != nil {
			return bobj.DeleteBucket([]byte(subb[0]))
		}

		return nil
	})
}

// Exists returns true if a bucket or sub-bucket exists in a database.
func Exists(db *bbolt.DB, buck string, subb ...string) (bool, error) {
	var ok bool

	return ok, db.View(func(tx *bbolt.Tx) error {
		if len(subb) == 0 {
			ok = tx.Bucket([]byte(buck)) != nil
			return nil
		}

		if bobj := tx.Bucket([]byte(buck)); bobj != nil {
			ok = bobj.Bucket([]byte(subb[0])) != nil
		}

		return nil
	})
}

// Get returns the pairs of an existing sub-bucket from a database.
func Get(db *bbolt.DB, buck, subb string) (map[string]string, error) {
	var pairs = make(map[string]string)

	return pairs, db.View(func(tx *bbolt.Tx) error {
		if bobj := tx.Bucket([]byte(buck)); bobj != nil {
			if sobj := bobj.Bucket([]byte(subb)); sobj != nil {
				return sobj.ForEach(func(attr []byte, data []byte) error {
					pairs[string(attr)] = string(data)
					return nil
				})
			}
		}

		return nil
	})
}

// List returns the sub-bucket names of an existing bucket from a database.
func List(db *bbolt.DB, buck string) ([]string, error) {
	var subbs []string

	return subbs, db.View(func(tx *bbolt.Tx) error {
		if bobj := tx.Bucket([]byte(buck)); bobj != nil {
			return bobj.ForEachBucket(func(subb []byte) error {
				subbs = append(subbs, string(subb))
				return nil
			})
		}

		return nil
	})
}

// Set sets the pairs in a new or existing sub-bucket in a database.
func Set(db *bbolt.DB, buck, subb string, pairs map[string]string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bobj, err := tx.CreateBucketIfNotExists([]byte(buck))
		if err != nil {
			return err
		}

		sobj, err := bobj.CreateBucketIfNotExists([]byte(subb))
		if err != nil {
			return err
		}

		for attr, data := range pairs {
			if err := sobj.Put([]byte(attr), []byte(data)); err != nil {
				return err
			}
		}

		return nil
	})
}
