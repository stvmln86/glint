package file

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/glint/glint/tools/path"
	"github.com/stvmln86/glint/glint/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	dest := filepath.Join(t.TempDir(), "test.extn")

	// success
	err := Create(dest, "Body.\n", 0777)
	test.AssertFile(t, dest, "Body.\n")
	assert.NoError(t, err)

	// error - already exists
	err = Create(dest, "Body.\n", 0777)
	assert.EqualError(t, err, `cannot create file "test.extn" - already exists`)
}

func TestDelete(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := path.Reextn(orig, ".trash")

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// error - does not exist
	err = Delete("/nope.extn")
	assert.EqualError(t, err, `cannot delete file "nope.extn" - does not exist`)
}

func TestExists(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success - true
	ok := Exists(orig)
	assert.True(t, ok)

	// success - false
	ok = Exists("/nope.extn")
	assert.False(t, ok)
}

func TestGlob(t *testing.T) {
	// setup
	dire := test.MockDire(t)

	// success
	origs := Glob(dire, ".extn")
	assert.Len(t, origs, 2)
	assert.Contains(t, origs, filepath.Join(dire, "alpha.extn"))
	assert.Contains(t, origs, filepath.Join(dire, "bravo.extn"))
}

func TestRead(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, "Alpha.\n", body)
	assert.NoError(t, err)

	// error - does not exist
	body, err = Read("/nope.extn")
	assert.Empty(t, body)
	assert.EqualError(t, err, `cannot read file "nope.extn" - does not exist`)
}

func TestRename(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := path.Rename(orig, "delta")

	// success
	err := Rename(orig, "delta")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// error - does not exist
	err = Rename("/nope.extn", dest)
	assert.EqualError(t, err, `cannot rename file "nope.extn" - does not exist`)
}

func TestSearch(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success - true
	ok, err := Search(orig, "ALPHA")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - false
	ok, err = Search(orig, "NOPE")
	assert.False(t, ok)
	assert.NoError(t, err)

	// error - does not exist
	ok, err = Search("/nope.extn", "NOPE")
	assert.False(t, ok)
	assert.EqualError(t, err, `cannot search file "nope.extn" - does not exist`)
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Update(orig, "Body.\n", 0777)
	test.AssertFile(t, orig, "Body.\n")
	assert.NoError(t, err)

	// error - does not exist
	err = Update("/nope.extn", "Body.\n", 0777)
	assert.EqualError(t, err, `cannot update file "nope.extn" - does not exist`)
}
