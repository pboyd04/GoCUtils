package strings

import (
	"testing"

	s "github.com/pboyd04/GoCUtils/pkg/strings"
)

func TestJustInt(t *testing.T) {
	result, str := s.Strtol("1", 10)
	if result != 1 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 1)
	}
	if len(str) != 0 {
		t.Errorf("Got an unexpected remainder, got %s", str)
	}
}

func TestJustIntWJunk(t *testing.T) {
	result, str := s.Strtol("1a", 10)
	if result != 1 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 1)
	}
	if str != "a" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "a")
	}
}

func TestMaxInt(t *testing.T) {
	result, str := s.Strtol("2147483647", 10)
	if result != 2147483647 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 2147483647)
	}
	if len(str) != 0 {
		t.Errorf("Got an unexpected remainder, got %s", str)
	}
}

func TestMaxIntWJunk(t *testing.T) {
	result, str := s.Strtol("2147483647?l", 10)
	if result != 2147483647 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 2147483647)
	}
	if str != "?l" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "?l")
	}
}

func TestMinInt(t *testing.T) {
	result, str := s.Strtol("-2147483646", 10)
	if result != -2147483646 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, -2147483646)
	}
	if len(str) != 0 {
		t.Errorf("Got an unexpected remainder, got %s", str)
	}
}

func TestMinIntWJunk(t *testing.T) {
	result, str := s.Strtol("-2147483646---", 10)
	if result != -2147483646 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, -2147483646)
	}
	if str != "---" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "---")
	}
}

func TestBaseDetection(t *testing.T) {
	result, str := s.Strtol("10", 0)
	if result != 10 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 10)
	}
	if str != "" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "")
	}

	result, str = s.Strtol("010", 0)
	if result != 8 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 8)
	}
	if str != "" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "")
	}

	result, str = s.Strtol("0x10", 0)
	if result != 16 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 16)
	}
	if str != "" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "")
	}
}

func TestBaseDetectionWJunk(t *testing.T) {
	result, str := s.Strtol("10a", 0)
	if result != 10 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 10)
	}
	if str != "a" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "a")
	}

	result, str = s.Strtol("010-", 0)
	if result != 8 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 8)
	}
	if str != "-" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "-")
	}

	result, str = s.Strtol("0x10x", 0)
	if result != 16 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 16)
	}
	if str != "x" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "x")
	}
}
