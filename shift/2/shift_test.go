package shift_test

import (
	"bytes"
	"testing"

	"github.com/bitfield/shift"
)

func TestEncipherTransformsHALToIBM(t *testing.T) {
	t.Parallel()
	input := []byte("HAL")
	want := []byte("IBM")
	got := shift.Encipher(input)
	if !bytes.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestEncipherTransformsADDToBEE(t *testing.T) {
	t.Parallel()
	input := []byte("ADD")
	want := []byte("BEE")
	got := shift.Encipher(input)
	if !bytes.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}
