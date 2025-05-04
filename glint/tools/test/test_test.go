package test

import (
	"errors"
	"testing"
)

func TestAssertDire(t *testing.T) {
	// setup
	dire := MockDire(t)

	// success
	AssertDire(t, dire, MockFiles)
}

func TestAssertErr(t *testing.T) {
	// setup
	err := errors.New(`test "error"`)

	// success
	AssertErr(t, err, "%s %q")
}

func TestAssertFile(t *testing.T) {
	// setup
	orig := MockFile(t, "alpha.extn")

	// success
	AssertFile(t, orig, "Alpha.\n")
}

func TestMockDire(t *testing.T) {
	// success
	dire := MockDire(t)
	AssertDire(t, dire, MockFiles)
}

func TestMockFile(t *testing.T) {
	// success
	orig := MockFile(t, "alpha.extn")
	AssertFile(t, orig, "Alpha.\n")
}
