package cipher_test

import (
	"cipher"
	"testing"
)

func TestShiftTransformsAToB(t *testing.T) {
	t.Parallel()
	input := 'A'
	want := 'B'
	got := cipher.Shift(input)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
