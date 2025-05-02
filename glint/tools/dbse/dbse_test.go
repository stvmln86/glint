package dbse

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/tools/test"
)

func TestOpen(t *testing.T) {
	// success
	db, err := Open(":memory:", "pragma foreign_keys = on")
	assert.NotNil(t, db)
	test.AssertBool(t, db, "pragma foreign_keys", true)
	assert.NoError(t, err)
}

func TestInit(t *testing.T) {
	// setup
	var size int
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success - first run
	okay, err := Init(db, "create table Test (a)")
	assert.True(t, okay)
	assert.NoError(t, err)

	// success - schema executed
	db.Get(&size, sCount)
	assert.Equal(t, 1, size)

	// success - second run
	okay, err = Init(db, "create table Test (a)")
	assert.False(t, okay)
	assert.NoError(t, err)

	// success - schema not executed
	db.Get(&size, sCount)
	assert.Equal(t, 1, size)
}
