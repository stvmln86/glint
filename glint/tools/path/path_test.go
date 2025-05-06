package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	// success
	base := Base("/dire/slug.extn")
	assert.Equal(t, "slug.extn", base)
}

func TestDire(t *testing.T) {
	// success
	dire := Dire("/dire/slug.extn")
	assert.Equal(t, "/dire", dire)
}

func TestExtn(t *testing.T) {
	// success - real extension
	extn := Extn("/dire/slug.extn")
	assert.Equal(t, ".extn", extn)

	// success - bare extension
	extn = Extn("/dire/slug.")
	assert.Equal(t, ".", extn)

	// success - no extension
	extn = Extn("/dire/slug")
	assert.Equal(t, "", extn)
}

func TestJoin(t *testing.T) {
	// success
	dest := Join("/dire", "slug", ".extn")
	assert.Equal(t, "/dire/slug.extn", dest)
}

func TestMatch(t *testing.T) {
	// success - true
	ok := Match("/dire/slug.extn", "SLUG")
	assert.True(t, ok)

	// success - false
	ok = Match("/dire/slug.extn", "NOPE")
	assert.False(t, ok)
}

func TestReslug(t *testing.T) {
	// success
	dest := Reslug("/dire/slug.extn", "test")
	assert.Equal(t, "/dire/test.extn", dest)
}

func TestSlug(t *testing.T) {
	// success
	slug := Slug("/dire/slug.extn")
	assert.Equal(t, "slug", slug)
}
