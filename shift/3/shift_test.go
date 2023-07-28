package shift_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bitfield/shift"
)

func TestEncipherTransforms(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		input, want []byte
	}{
		{
			input: []byte("HAL"),
			want:  []byte("IBM"),
		},
		{
			input: []byte("ADD"),
			want:  []byte("BEE"),
		},
		{
			input: []byte("ANA"),
			want:  []byte("BOB"),
		},
		{
			input: []byte("INKS"),
			want:  []byte("JOLT"),
		},
		{
			input: []byte("ADMIX"),
			want:  []byte("BENJY"),
		},
		{
			input: []byte{0, 1, 2, 3, 255},
			want:  []byte{1, 2, 3, 4, 0},
		},
	}
	for _, tc := range tcs {
		name := fmt.Sprintf("%s to %s", tc.input, tc.want)
		t.Run(name, func(t *testing.T) {
			got := shift.Encipher(tc.input)
			if !bytes.Equal(tc.want, got) {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}
