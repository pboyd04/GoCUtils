package strings

import (
	"testing"

	s "github.com/pboyd04/GoCUtils/pkg/strings"
)

func TestStrtollJustInt(t *testing.T) {
	result, str := s.Strtoll("1", 10)
	if result != 1 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 1)
	}
	if len(str) != 0 {
		t.Errorf("Got an unexpected remainder, got %s", str)
	}
}

func TestStrtollJustIntWJunk(t *testing.T) {
	result, str := s.Strtoll("1a", 10)
	if result != 1 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 1)
	}
	if str != "a" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "a")
	}
}

func TestStrtollMaxInt(t *testing.T) {
	result, str := s.Strtoll("9223372036854775807", 10)
	const max int64 = 9223372036854775807
	if result != max {
		t.Errorf("Result was incorrect, got %d, want %d.", result, max)
	}
	if len(str) != 0 {
		t.Errorf("Got an unexpected remainder, got %s", str)
	}
}

func TestStrtollMaxIntWJunk(t *testing.T) {
	result, str := s.Strtoll("9223372036854775807?l", 10)
	const max int64 = 9223372036854775807
	if result != max {
		t.Errorf("Result was incorrect, got %d, want %d.", result, max)
	}
	if str != "?l" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "?l")
	}
}

func TestStrtollMinInt(t *testing.T) {
	result, str := s.Strtoll("-9223372036854775806", 10)
	if result != -9223372036854775806 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, -9223372036854775806)
	}
	if len(str) != 0 {
		t.Errorf("Got an unexpected remainder, got %s", str)
	}
}

func TestStrtollMinIntWJunk(t *testing.T) {
	result, str := s.Strtoll("-9223372036854775806---", 10)
	if result != -9223372036854775806 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, -9223372036854775806)
	}
	if str != "---" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "---")
	}
}

func TestStrtollBaseDetection(t *testing.T) {
	result, str := s.Strtoll("10", 0)
	if result != 10 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 10)
	}
	if str != "" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "")
	}

	result, str = s.Strtoll("010", 0)
	if result != 8 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 8)
	}
	if str != "" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "")
	}

	result, str = s.Strtoll("0x10", 0)
	if result != 16 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 16)
	}
	if str != "" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "")
	}
}

func TestStrtollBaseDetectionWJunk(t *testing.T) {
	result, str := s.Strtoll("10a", 0)
	if result != 10 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 10)
	}
	if str != "a" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "a")
	}

	result, str = s.Strtoll("010-", 0)
	if result != 8 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 8)
	}
	if str != "-" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "-")
	}

	result, str = s.Strtoll("0x10x", 0)
	if result != 16 {
		t.Errorf("Result was incorrect, got %d, want %d.", result, 16)
	}
	if str != "x" {
		t.Errorf("Remainder was incorrect, got %s, want %s.", str, "x")
	}
}
