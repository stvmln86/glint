// Package test implements unit testing helper functions and data.
package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/tools/sqls"
)

// MockInserts is a set of mock database inserts for unit testing.
const MockInserts = `
	insert into Notes (init, name) values (1000, 'alpha');
	insert into Pages (init, note, body) values (1000, 1, 'Alpha one.\n');
	insert into Pages (init, note, body) values (1100, 1, 'Alpha two.\n');

	insert into Notes (init, name) values (2000, 'bravo');
	insert into Pages (init, note, body) values (2000, 1, 'Bravo one.\n');
`

// AssertBool asserts a database query result is equal to an integer.
func AssertBool(t *testing.T, db *sqlx.DB, code string, want bool) {
	var data bool
	err := db.Get(&data, code)
	assert.Equal(t, want, data)
	assert.NoError(t, err)
}

// AssertInt asserts a database query result is equal to an integer.
func AssertInt(t *testing.T, db *sqlx.DB, code string, want int) {
	var data int
	err := db.Get(&data, code)
	assert.Equal(t, want, data)
	assert.NoError(t, err)
}

// AssertStr asserts a database query result is equal to a string.
func AssertStr(t *testing.T, db *sqlx.DB, code, want string) {
	var data string
	err := db.Get(&data, code)
	assert.Equal(t, want, data)
	assert.NoError(t, err)
}

// DB returns a new in-memory database containing mock data.
func DB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(sqls.Pragma + sqls.Schema + MockInserts)
	return db
}
