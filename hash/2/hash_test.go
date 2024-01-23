package hash_test

import (
	"bytes"
	"testing"

	"github.com/bitfield/hash"
)

func TestSumHashReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	input := []byte("I love you, Bob")
	want := []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x04, 0xfb,
	}
	got := hash.SumHash(input)
	if !bytes.Equal(want, got) {
		t.Errorf("%s: want %x, got %x", input, want, got)
	}
}
