// Package dbse implements database handling functions.
package dbse

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const sCount = "select count(*) from SQLITE_SCHEMA"

// Open returns a new database connection with executed pragma.
func Open(path, prag string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("cannot open database - %w", err)
	}

	if _, err := db.Exec(prag); err != nil {
		return nil, fmt.Errorf("cannot query database - %w", err)
	}

	return db, nil
}

// Init executes schema on an undefined database and returns true if the schema was
// executed.
func Init(db *sqlx.DB, schm string) (bool, error) {
	var size int
	if err := db.Get(&size, sCount); err != nil {
		return false, fmt.Errorf("cannot query database - %w", err)
	}

	if size == 0 {
		if _, err := db.Exec(schm); err != nil {
			return false, fmt.Errorf("cannot query database - %w", err)
		}

		return true, nil
	}

	return false, nil
}
