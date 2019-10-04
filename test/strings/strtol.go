package strings

import (
	"testing"

	. "github.com/pboyd04/GoCUtils/strings"
)

func TestJustInt(t *testing.T) {
	result, str := Strtol("1")
	if result != 1 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 1)
	}
	if len(str) != 0 {
		t.Errorf("Got an unexpected remainder, got %s", str)
	}
}
