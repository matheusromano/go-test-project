package strings

import (
	"testing"
)

func TestCompare(t *testing.T) {
	got := StringCompare("lele", "lele")
	expected := 0

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}

}

func TestContains(t *testing.T) {
	got := StringContains("age of impire", "age")
	expected := true

	if got != expected {
		t.Errorf("expected %t but got %t", expected, got)
	}
}
