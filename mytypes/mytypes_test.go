package mytypes_test

import (
	"mytypes"
	"strings"
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

func TestStringsBuilder(t *testing.T) {
	t.Parallel()

	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := sb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("expected %d with string '%q', got %d", wantLen, sb.String(), gotLen)
	}
}
