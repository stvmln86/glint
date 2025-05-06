package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TextExtn(t *testing.T) {
	// success - with dot
	extn := Extn("\t.EXTN\n")
	assert.Equal(t, ".extn", extn)

	// success - without dot
	extn = Extn("\tEXTN\n")
	assert.Equal(t, ".extn", extn)
}

func TestPath(t *testing.T) {
	// success
	path := Path("\t/././dire/name.extn\n")
	assert.Equal(t, "/dire/name.extn", path)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME 123 !!!\n")
	assert.Equal(t, "name-123", name)
}
