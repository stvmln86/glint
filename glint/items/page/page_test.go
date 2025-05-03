package page

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/tools/bolt"
	"github.com/stvmln86/glint/glint/tools/neat"
	"github.com/stvmln86/glint/glint/tools/test"
)

func xPage(t *testing.T) *Page {
	db := test.MockDB(t)
	pairs := test.MockData["alpha"]["1000"]
	return &Page{db, "alpha", pairs["body"], pairs["hash"], "1000"}
}

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB(t)
	hash := neat.Hash("Alpha (test).\n")
	init := neat.Unix(time.Now())

	// success
	page, err := Create(db, "alpha", "Alpha (test).\n")
	assert.NotNil(t, page.DB)
	assert.Equal(t, "alpha", page.Name)
	assert.Equal(t, "Alpha (test).\n", page.Body)
	assert.Equal(t, hash, page.Hash)
	assert.Equal(t, init, page.Init)
	assert.NoError(t, err)

	// success - check database
	pairs, err := bolt.Get(db, "alpha", init)
	assert.Equal(t, "Alpha (test).\n", pairs["body"])
	assert.NoError(t, err)

	// error - does not exist
	page, err = Create(db, "nope", "Body.\n")
	assert.Nil(t, page)
	assert.EqualError(t, err, "cannot append note nope - does not exist")
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	page, err := Get(db, "alpha", "1000")
	assert.NotNil(t, page.DB)
	assert.Equal(t, "alpha", page.Name)
	assert.Equal(t, "Alpha (old).\n", page.Body)
	assert.Equal(t, test.MockData["alpha"]["1000"]["hash"], page.Hash)
	assert.Equal(t, "1000", page.Init)
	assert.NoError(t, err)

	// error - does not exist
	page, err = Get(db, "alpha", "nope")
	assert.Nil(t, page)
	assert.EqualError(t, err, "cannot get page alpha/nope - does not exist")
}

func TestDelete(t *testing.T) {
	// setup
	page := xPage(t)

	// success
	err := page.Delete()
	assert.NoError(t, err)

	// success - check database
	ok, err := bolt.Exists(page.DB, "alpha", "1000")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	page := xPage(t)

	// success
	ok, err := page.Exists()
	assert.True(t, ok)
	assert.NoError(t, err)
}
