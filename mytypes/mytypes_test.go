package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()

	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()
	if want != got {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestStrLen(t *testing.T) {
	t.Parallel()

	input := mytypes.MyStr("Hello")
	want := 5
	got := input.StrLen()
	if want != got {
		t.Errorf("expected %d with string '%s', got %d", want, input, got)
	}
}
