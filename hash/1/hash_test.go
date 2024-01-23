package hash_test

import (
	"bytes"
	"testing"

	"github.com/bitfield/hash"
)

func TestLenHashReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	input := []byte("I love you, Bob")
	want := []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x0f,
	}
	got := hash.LenHash(input)
	if !bytes.Equal(want, got) {
		t.Errorf("%s: want %x, got %x", input, want, got)
	}
}
