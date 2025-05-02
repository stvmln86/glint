package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/tools/sqls"
)

func TestMockInserts(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(sqls.Schema + MockInserts)
	assert.NoError(t, err)
}

func TestAssertBool(t *testing.T) {
	// setup
	db := DB()

	// success
	AssertBool(t, db, "select exists(select * from Notes)", true)
}

func TestAssertInt(t *testing.T) {
	// setup
	db := DB()

	// success
	AssertInt(t, db, "select n_id from Notes where name='alpha'", 1)
}

func TestAssertStr(t *testing.T) {
	// setup
	db := DB()

	// success
	AssertStr(t, db, "select name from Notes where n_id=1", "alpha")
}

func TestDB(t *testing.T) {
	// success
	db := DB()
	assert.NotNil(t, db)
}
