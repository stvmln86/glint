package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestHash(t *testing.T) {
	// success
	hash := Hash("data")
	assert.Equal(t, "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447", hash)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME\n")
	assert.Equal(t, "name", name)
}

func TestPairs(t *testing.T) {
	// success
	pairs := Pairs("\tBody.\n")
	assert.Equal(t, map[string]string{
		"body": "Body.\n",
		"hash": "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447",
	}, pairs)
}

func TestTime(t *testing.T) {
	// setup
	want := time.UnixMilli(1000)

	// success
	tobj := Time("1000")
	assert.Equal(t, want, tobj)
}

func TestUnix(t *testing.T) {
	// setup
	tobj := time.UnixMilli(1000)

	// success
	unix := Unix(tobj)
	assert.Equal(t, "1000", unix)
}
