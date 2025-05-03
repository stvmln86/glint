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
	assert.Equal(t, "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7", hash)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME\n")
	assert.Equal(t, "name", name)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(1000, 0)

	// success
	tobj := Time("1000")
	assert.Equal(t, want, tobj)
}

func TestUnix(t *testing.T) {
	// setup
	tobj := time.Unix(1000, 0)

	// success
	unix := Unix(tobj)
	assert.Equal(t, "1000", unix)
}
